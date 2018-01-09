export default class SendMessageForm extends React.Component {
    onSubmit = (e) => {
        e.preventDeafult();
        this.props.submitMessage(this.input)
    }

    render () {
        return <div className='row'>
            <div className='col-12' onSubmit={this.onSubmit}>
                <form>
                    <input value={this.input.value} ref={(input => this.input = input)}/>
                    <input type='submit' value='Send'/>
                </form>
            </div>
        </div>
    }
}