var React = require('react');
var request = require('request');
var $ = require('jquery');
var _ = require('lodash');
var Display = require('./Display.jsx');

var Profile = React.createClass({
	getInitialState: function(){
		return{tickets: []};
	},
	componentWillMount: function(){
		request({
			url: 'http://www.ruexploring.com:80/GetUserValidTickets',
			method: 'POST',
			json: {
				ID: this.props.userID
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
                if(body == "Active Coupons Does Not Exist"){
					this.setState({tickets: body});
				}else{
					var objects = _.cloneDeep(body.coupons);
					this.setState({tickets: objects});
				}
			}
		}.bind(this));
	},
	useCoupon: function(id){
		request({
			url: 'http://www.ruexploring.com:80/UseTicket',
			method: 'POST',
			json: {
				ID: this.props.userID,
				COUPON_ID: id
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
			}
		}.bind(this));

		this.forceUpdate();
	},
	render: function(){
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Current Coupons:</strong></h3>
					</div>
				</div>
				<Display tickets={this.state.tickets} useCoupon={this.useCoupon}/>
			</div>
		);
	}
});

module.exports = Profile;
