var React = require('react');
var NavBar = require('./NavBar.jsx');
var Page = require('./Page.jsx');
var Footer = require('./Footer.jsx');

var Window = React.createClass({
    getInitialState: function(){
        return {state: "Home"};
    },
    setPageState: function(state){
		this.setState({state: state});
	},
    render: function(){
        return (
            <div>
                <NavBar stateCallback={this.setPageState}></NavBar>
                <Page stateCallback={this.setPageState} state={this.state.state}></Page>
				<br></br>
				<Footer></Footer>
			</div>
        );
    }
});

React.render(<Window/>, document.getElementById('app'));
