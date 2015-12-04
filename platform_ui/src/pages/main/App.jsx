var React = require('react');
var User = require('./User.jsx');
var Business = require('./Business.jsx');

var App = React.createClass({
	getInitialState: function(){
		return({UserType: "User"});
	},
	stateCallback: function(state){
		this.props.stateCallback(state);
	},
	render: function(){
		var renderUserType;
		switch(this.state.UserType){
			case "User":
				renderUserType = <User stateCallback={this.props.stateCallback}/>
				break;
			case "Business":
				renderUserType = <Business stateCallback={this.props.stateCallback}/>
				break;
			default:
				console.log("Error UserType Undefined");
		}

		return(
			<div className="container">
				{renderUserType}
			</div>
		);
	}
});

module.exports = App;
