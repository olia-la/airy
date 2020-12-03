import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import _, {connect, ConnectedProps} from 'react-redux';
// import ReactGA from 'react-ga';
import {AccessibleSVG} from '../../../components/AccessibleSVG';
import IconChannel from '../../../components/IconChannel';

import styles from './index.module.scss';

import {formatTimeOfMessage} from '../../../services/format/date';
// import {toggleConversationStatus} from 'airy-client/actions/conversations';
// import {toggleState} from '../../../services/conversation';
import {MESSENGER_CONVERSATIONS_ROUTE} from '../../../routes/routes';
import {Conversation} from '../../../model/Conversation';
import {Message} from '../../../model/Message';
import {StateModel} from '../../../reducers';

const FormattedMessage = ({message}: {message: Message}) => {
  if (message && message.content) {
    return <>{message.content.text}</>;
  }
  return <div />;
};

const mapStateToProps = (state: StateModel) => {
  return {
    // contacts: state.data.contacts.items,
    // channels: state.data.channels.data,
  };
};

const connector = connect(mapStateToProps, null);

type Props = {
  conversation: Conversation & {unread_message_count: number; message?: Message};
  active: boolean;
  style: any;
} & ConnectedProps<typeof connector>;
class ConversationListItem extends Component<Props, null> {
  render() {
    const {conversation, active, style} = this.props;

    const participant = {
      avatar_url: 'https://microhealth.com/assets/images/illustrations/personal-user-illustration-@2x.png',
      first_name: 'Random',
      last_name: 'User',
      display_name: 'Random User',
      info: {},
    };
    const unread = conversation.unread_message_count > 0;

    return (
      <div className={styles.clickableListItem} style={style}>
        <Link to={`${MESSENGER_CONVERSATIONS_ROUTE}/${conversation.id}`}>
          <div
            className={`${active ? styles.containerListItemActive : styles.containerListItem} ${
              unread ? styles.unread : ''
            }`}>
            <div
              className={styles.profileImage}
              style={{backgroundImage: `url(${participant && participant.avatar_url})`}}
            />
            <div className={styles.contactDetails}>
              <div className={styles.topRow}>
                <div className={`${styles.profileName} ${unread ? styles.unread : ''}`}>
                  {participant && participant.display_name}
                </div>
                <div className={styles.statusIcon}>
                  {/* <ConversationStatus t={t} onClick={this.toggleConversationState} status={conversation.state} /> */}
                </div>
              </div>
              <div className={`${styles.contactLastMessage} ${unread ? styles.unread : ''}`}>
                <FormattedMessage message={conversation.last_message} />
              </div>
              <div className={styles.bottomRow}>
                <div className={styles.source}>
                  <IconChannel channel={conversation.channel} avatar={true} name={true} />
                </div>
                <div className={styles.contactLastMessageDate}>{formatTimeOfMessage(conversation.last_message)}</div>
              </div>
            </div>
          </div>
        </Link>
      </div>
    );
  }
}

export default connector(ConversationListItem);