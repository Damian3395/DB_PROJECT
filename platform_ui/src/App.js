var React = require('react');
var NavBar = require('./NavBar.jsx');
var Page = require('./Page.jsx');

var Window = React.createClass({
    getInitialState: function(){
        return {state: "Home"};
    },
    setState: function(state){
        this.state.state= state;
        console.log("New State %s", this.state.state);
        this.forceUpdate();
    },
    render: function(){
        return (
            <div>
                <NavBar stateCallback={this.setState}></NavBar>
                <Page state={this.state.state}></Page>
            </div>
        );
    }
});

React.render(<Window/>, document.getElementById('app'));
