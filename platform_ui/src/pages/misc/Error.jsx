var React = require('react');

var Error = React.createClass({
	getInitialState: function(){
		return({error: this.props.error});
	},
	componentWillReceiveProps: function(nextProps){
		this.setState({error: nextProps.error});
	},
	render: function(){
		if(this.state.error !== ""){
			return(
				<div className="row text-center">
					<h1 className="label label-danger">{this.state.error}</h1>
				</div>
			);
		}else{
			return false;
		}
	}
});

module.exports = Error;
