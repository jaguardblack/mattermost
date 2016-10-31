// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
import React from 'react';
import {FormattedMessage} from 'react-intl';

export default class NewMessageIndicator extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            visible: false,
            rendered: false
        }
    }
    componentDidMount() {
        this._onTransition = this.onTransition.bind(this);
        this.refs.indicator.addEventListener("transitionend", this._onTransition);
    }
    componentWillUnmount() {
        this.refs.indicator.removeEventListener("transitionend", this._onTransition);
    }
    componentWillReceiveProps(nextProps) {
      if(nextProps.newMessages > 0) {
        this.setState({rendered: true}, () => {
          this.setState({visible: true});
        })
      } else {
        this.setState({visible: false})
      }
    }
    render() {
        let className = 'nav-pills__unread-indicator-bottom';
        if (this.state.visible > 0) {
            className += ' visible';
        }
        if (!this.state.rendered) {
            className += ' disabled';
        }
        return (
            <div
                className={className}
                onClick={this.props.onClick}
                ref='indicator'
            >
                <span>
                    <i
                        className='fa fa-arrow-circle-o-down'
                    />
                    <FormattedMessage
                        id='posts_view_newMsgBelow'
                        defaultMessage='{count} new {count, plural, one {message} other {messages}} below'
                        values={{count: this.props.newMessages}}
                    />
                </span>
            </div>
        );
    }
    onTransition() {
        this.setState({rendered: this.state.visible })
    }
}
NewMessageIndicator.defaultProps = {
    newMessages: 0
};

NewMessageIndicator.propTypes = {
    onClick: React.PropTypes.func.isRequired,
    newMessages: React.PropTypes.number
};
