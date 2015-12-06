var React = require('react');
var User = require('./User.jsx');
var Business = require('./Business.jsx');

var App = React.createClass({
	stateCallback: function(state){
		this.props.stateCallback(state);
	},
	render: function(){
		var renderUserType;
		switch(this.props.type){
			case "User":
				renderUserType = <User stateCallback={this.props.stateCallback} userID={this.props.userID}/>
				break;
			case "Business":
				renderUserType = <Business stateCallback={this.props.stateCallback} userID={this.props.userID}/>
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
