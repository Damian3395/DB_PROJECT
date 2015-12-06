var React = require('react');
var request = require('request');
var $ = require('jquery');
var _ = require('lodash');
var Coupon = require('./Coupon.jsx');

var Profile = React.createClass({
	getInitialState: function(){
		return{};
	},
	componentWillMount: function(){
		request({
			url: 'http://localhost:8080/GetCurrentCoupons',
			method: 'POST',
			json: {
				ID: "USERID",
				TOKEN: "TOKEN"
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
			}
		}.bind(this));

		request({
			url: 'http://localhost:8080/GetUsedCoupons',
			method: 'POST',
			json: {
				ID: "USERID",
				TOKEN: "TOKEN"
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
			}
		}.bind(this));
		console.log("Update Page State");
	},
	render: function(){
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Current Coupons:</strong></h3>
					</div>
				</div>
				
				<Coupon type="In-Use"/>

				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Used Coupons:</strong></h3>
					</div>
				</div>
				
				<Coupon type="Used"/>
			</div>
		);
	}
});

module.exports = Profile;
