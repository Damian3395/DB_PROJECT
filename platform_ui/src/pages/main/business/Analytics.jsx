var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var request = require('request');
var Charts = require('./Charts.jsx');

var Analytics = React.createClass({
	getIntialState: function(){
		return{coupons: [], Data: []};
	},
	componentWillMount: function(){
		request({
			url: 'http://localhost:8080/GetExpiredCoupons',
			method: 'POST',
			json: {
				ID: this.props.userID,
			}
		}, function(error, response, body){
			console.log(body)
			if(error){
				console.log(error);
			}else if(body == "Error"){
				this.setState({coupons: body});	
			}else{
				console.log(response.statusCode, body);
				this.setState({coupons: body.coupons});
			}
		}.bind(this));
	},
	DisplayChart: function(){
		var id = $('#COUPON_ID').val()
		if(id != "DNE"){
			request({
				url: 'http://localhost:8080/GetCouponAnalytics',
				method: 'POST',
				json: {
					COUPON_ID: id,
				}
			}, function(error, response, body){
				if(error){
					console.log(error);
				}else{
					console.log(response.statusCode, body);
					this.setState({Data: body});
				}
			}.bind(this));
		}
	},
	render: function(){
		var DISPLAY;
		if(this.state.coupons == "Expired Coupons Does Not Exist"){
			DISPLAY = <option id="DNE">Expired Coupons Does Not Exist</option>
		}else{
			DISPLAY = this.state.coupons.map(function(coupon) {
				return <option id={coupon.COUPON_ID}>{coupon.TYPE} {coupon.MONTH}-{coupon.DAY}-{coupon.YEAR}</option>;
			});
		}
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3><strong>Analytics</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-7">
						<div className="input-group">
							<span className="input-group-addon">Select Coupon</span>
							<select className="form-control" id="COUPON_ID">
								{DISPLAY}				
							</select>
						</div>	
					</div>
					<div className="col-md-offset-1 col-md-4">
						<button type="button"
							className="btn btn-success btn-block"
							onClick={this.DisplayChart}>
								Display
						</button>
					</div>
				</div>
				<br/>
				<div className="row">
					<Charts Data={this.state.Data}/>
				</div>
			</div>
		);
	}
});

module.exports = Analytics;
