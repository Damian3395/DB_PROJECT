var React = require('react');
var request = require('request');
var Coupon = require('./Coupon.jsx');

var Profile = React.createClass({
	getInitialState: function(){
		return{};
	},
	componentWillMount: function(){
		console.log("Get User Coupon Information From DataBase");	
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
