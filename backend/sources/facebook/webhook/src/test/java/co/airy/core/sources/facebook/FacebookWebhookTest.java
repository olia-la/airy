package co.airy.core.sources.facebook;

import co.airy.kafka.schema.Topic;
import co.airy.kafka.schema.source.SourceFacebookEvents;
import co.airy.kafka.test.KafkaTestHelper;
import co.airy.kafka.test.junit.SharedKafkaTestResource;
import co.airy.spring.core.AirySpringBootApplication;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.api.extension.RegisterExtension;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.Matchers.hasSize;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest(
        properties = {
                "facebook.webhook-secret=theansweris42"
        },
        webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT, classes = AirySpringBootApplication.class)
@AutoConfigureMockMvc
@ExtendWith(SpringExtension.class)
class FacebookWebhookTest {
    @RegisterExtension
    public static final SharedKafkaTestResource sharedKafkaTestResource = new SharedKafkaTestResource();

    private static KafkaTestHelper kafkaTestHelper;

    private static final Topic sourceFacebookEvents = new SourceFacebookEvents();

    @Autowired
    private MockMvc mvc;

    @BeforeAll
    static void beforeAll() throws Exception {
        kafkaTestHelper = new KafkaTestHelper(sharedKafkaTestResource, sourceFacebookEvents);

        kafkaTestHelper.beforeAll();
    }

    @Test
    void returns200() throws Exception {
        mvc.perform(post("/facebook")
                .content("whatever"))
                .andExpect(status().isOk());

        List<String> records = kafkaTestHelper.consumeValues(1, sourceFacebookEvents.name());

        assertThat(records, hasSize(1));
        assertEquals("whatever", records.get(0));
    }

    @AfterAll
    static void afterAll() throws Exception {
        kafkaTestHelper.afterAll();
    }
}
