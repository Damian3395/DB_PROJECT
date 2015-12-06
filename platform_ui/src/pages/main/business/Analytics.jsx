var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var request = require('request');

var Analytics = React.createClass({
	componentWillMount: function(){
		request({
			url: 'http://localhost:8080/GetExpiredCoupons',
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
				console.log("Update Page State");
			}
		}.bind(this));
	},
	render: function(){
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3><strong>Testing Analytics</strong></h3>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = Analytics;
