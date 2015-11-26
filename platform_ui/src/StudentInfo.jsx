var React = require('react');

var StudentInfo = React.createClass({
	render: function(){
		if(this.props.isStudent){
			return this.props.children;
		}else{
			return false;
		}
	}
});

module.exports = StudentInfo;
