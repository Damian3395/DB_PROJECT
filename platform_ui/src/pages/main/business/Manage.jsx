var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var request = require('request');
var Coupon = require('./Coupon.jsx');

var Manage = React.createClass({
	getInitialState: function(){
		return{student: false, c: []};
	},
	shouldComponentUpdate: function(nextProps, nextState) {
		return true;
	},
	componentWillMount: function(){
		request({
			url: 'http://www.ruexploring.com/GetActiveCoupons',
			method: 'POST',
			json: {
				ID: this.props.userID
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				if(body == "None"){
					console.log("No Active Coupons")
					this.setState({c: "No Active Coupons"});	
				}else{
					console.log(body.coupons);
					var objects = _.cloneDeep(body.coupons);
					this.setState({c: objects});
				}
			}
		}.bind(this));	
	},
	createCoupon: function(){
		var student_only = "False";
		if(this.state.student){
			student_only = "True"
		}
		request({
			url: 'http://localhost:8080/CreateCoupon',
			method: 'POST',
			json: {
				ID: this.props.userID,
				TYPE: $('#createType').val(),
				DAY: $('#day').val(),
				MONTH: $('#month').val(),
				YEAR: $('#year').val(),
				VALID: "TRUE",
				STUDENT: student_only
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
	removeCoupon: function(id){
		request({
			url: 'http://localhost:8080/RemoveCoupon',
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
	},
	isStudent: function(){
		this.setState({student: !this.state.student});
	},
	render: function(){
		var DISPLAY;
		if(this.state.c == "No Active Coupons"){
			DISPLAY = <div>NO ACTIVE COUPONS</div>
		}else{
			var COUPONS = this.state.c.map(function(coupon) {
				return <div className="row">
							<div className="col-md-8">
								<Coupon type={coupon.TYPE} day={coupon.DAY} month={coupon.MONTH} year={coupon.YEAR}/>	
							</div>
							<div clasName="col-md-4">
								<button type="button"
									className="btn btn-md btn-danger"
									onClick={this.removeCoupon.bind(this, coupon.COUPON_ID)}>
										Remove
								</button>
							</div>	
						</div>
				;
			}.bind(this));
			DISPLAY = <div>{COUPONS}</div>;
		}
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Create New Coupon</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-12">
						<div className="input-group">
							<span className="input-group-addon">Coupon Type: </span>
							<select className="form-control" id="createType">
								<option value="10% off">10% off</option>
								<option value="20% off">20% off</option>
								<option value="30% off">30% off</option>
								<option value="40% off">40% off</option>
								<option value="50% off">50% off</option>
								<option value="60% off">60% off</option>
								<option value="70% off">70% off</option>
							</select>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<h5>Expiration Date:</h5>
					</div>
				</div>
				<div className="row">
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Day:</span>
							<input type="number" className="form-control" id="day"/>
						</div>
					</div>
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Month:</span>
							<input type="number" className="form-control" id="month"/>
						</div>
					</div>
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Year:</span>
							<input type="number" className="form-control" id="year"/>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<div className="checkbox">
							<label><input type="checkbox" id="student" onClick={this.isStudent}/>For Students Only?</label>
						</div>
					</div>
					<div className="col-md-offset-4 col-md-4">
						<button type="button"
							className="btn btn-md btn-danger btn-block"
							onClick={this.createCoupon}>
								Create Coupon
						</button>
					</div>
				</div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Current Coupons</strong></h3>
					</div>
				</div>
				{DISPLAY}
			</div>
		);
	}
});

module.exports = Manage;
