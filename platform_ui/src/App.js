var React = require('react');
var NavBar = require('./NavBar.jsx');
var Page = require('./Page.jsx');
var Footer = require('./Footer.jsx');

var Window = React.createClass({
    getInitialState: function(){
        return {state: "Home", type: "", userID: ""};
    },
    setPageState: function(state, type, userID){
		this.setState({state: state, type: type, userID: userID});
	},
    render: function(){
        return (
            <div id="wrapper">
                <NavBar stateCallback={this.setPageState} state={this.state.state} type={this.state.type} userID={this.state.userID}></NavBar>
                <Page stateCallback={this.setPageState} state={this.state.state} type={this.state.type} userID={this.state.userID}></Page>
				<br></br>
				<Footer></Footer>
			</div>
        );
    }
});

React.render(<Window/>, document.getElementById('app'));
