import ReactDOM from 'react-dom';
import React from 'react';

export default class SendMessageForm extends React.Component {
    state = {
        input: ''
    }

    onChange = (e) => {
        this.setState({
            input: e.target.value
        })
    }

    onSubmit = (e) => {
        e.preventDeafult();
        this.props.submitMessage(this.state.input)
    }

    render () {
        return <div className='row'>
            <div className='col-12' onSubmit={this.onSubmit}>
                <form>
                    <input value={this.state.input} onChange={this.onChange} />
                    <input type='submit' value='Send'/>
                </form>
            </div>
        </div>
    }
}