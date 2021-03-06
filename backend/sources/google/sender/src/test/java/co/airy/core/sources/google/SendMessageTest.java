package co.airy.core.sources.google;

import co.airy.avro.communication.DeliveryState;
import co.airy.avro.communication.Message;
import co.airy.avro.communication.SenderType;
import co.airy.core.sources.google.model.SendMessagePayload;
import co.airy.core.sources.google.services.Api;
import co.airy.kafka.schema.Topic;
import co.airy.kafka.schema.application.ApplicationCommunicationMessages;
import co.airy.kafka.test.KafkaTestHelper;
import co.airy.kafka.test.junit.SharedKafkaTestResource;
import co.airy.spring.core.AirySpringBootApplication;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.api.extension.RegisterExtension;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.MockitoAnnotations;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import java.time.Instant;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

import static co.airy.test.Timing.retryOnException;
import static org.apache.kafka.streams.KafkaStreams.State.RUNNING;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.core.IsEqual.equalTo;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.doNothing;

@SpringBootTest(classes = AirySpringBootApplication.class)
@TestPropertySource(value = "classpath:test.properties")
@ExtendWith(SpringExtension.class)
class SendMessageTest {
    @RegisterExtension
    public static final SharedKafkaTestResource sharedKafkaTestResource = new SharedKafkaTestResource();
    private static KafkaTestHelper kafkaTestHelper;

    private static final Topic applicationCommunicationMessages = new ApplicationCommunicationMessages();

    @MockBean
    private Api api;

    @Autowired
    @InjectMocks
    private Sender worker;

    private static boolean streamInitialized = false;

    @BeforeAll
    static void beforeAll() throws Exception {
        kafkaTestHelper = new KafkaTestHelper(sharedKafkaTestResource, applicationCommunicationMessages);
        kafkaTestHelper.beforeAll();
    }

    @AfterAll
    static void afterAll() throws Exception {
        kafkaTestHelper.afterAll();
    }

    @BeforeEach
    void beforeEach() throws InterruptedException {
        MockitoAnnotations.initMocks(this);

        if (!streamInitialized) {
            retryOnException(() -> assertEquals(worker.getStreamState(), RUNNING), "Failed to reach RUNNING state.");
            streamInitialized = true;
        }
    }

    @Test
    void canSendMessageViaGoogleApi() throws Exception {
        final String conversationId = UUID.randomUUID().toString();
        final String messageId = UUID.randomUUID().toString();
        final String sourceConversationId = "source-conversation-id";
        final String channelId = UUID.randomUUID().toString();
        final String text = "Hello World";

        ArgumentCaptor<SendMessagePayload> payloadCaptor = ArgumentCaptor.forClass(SendMessagePayload.class);
        ArgumentCaptor<String> googleIdCaptor = ArgumentCaptor.forClass(String.class);
        doNothing().when(api).sendMessage(googleIdCaptor.capture(), payloadCaptor.capture());

        // Populate the source conversation id in the store
        kafkaTestHelper.produceRecord(
                new ProducerRecord<>(applicationCommunicationMessages.name(), UUID.randomUUID().toString(),
                        Message.newBuilder()
                                .setId("other-message-id")
                                .setSource("google")
                                .setSentAt(Instant.now().toEpochMilli())
                                .setSenderId(sourceConversationId)
                                .setSenderType(SenderType.SOURCE_CONTACT)
                                .setDeliveryState(DeliveryState.DELIVERED)
                                .setConversationId(conversationId)
                                .setChannelId(channelId)
                                .setContent("{\"text\":\"" + text + "\"}")
                                .build())
        );

        TimeUnit.SECONDS.sleep(5);

        kafkaTestHelper.produceRecord(new ProducerRecord<>(applicationCommunicationMessages.name(), messageId,
                Message.newBuilder()
                        .setId(messageId)
                        .setSentAt(Instant.now().toEpochMilli())
                        .setSenderId("user-id")
                        .setSenderType(SenderType.APP_USER)
                        .setDeliveryState(DeliveryState.PENDING)
                        .setConversationId(conversationId)
                        .setChannelId(channelId)
                        .setSource("google")
                        .setContent("{\"text\":\"" + text + "\"}")
                        .build())
        );

        retryOnException(() -> {
            final SendMessagePayload sendMessagePayload = payloadCaptor.getValue();

            assertThat(sendMessagePayload.getText(), equalTo(text));
            assertThat(sendMessagePayload.getMessageId(), equalTo(messageId));
            assertThat(googleIdCaptor.getValue(), equalTo(sourceConversationId));
        }, "google API was not called");
    }
}
