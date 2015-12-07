var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var request = require('request');
var Options = require('../../misc/Options.jsx');

var Settings = React.createClass({
	componentWillMount: function(){
		request({
			url: 'http://www.ruexploring.com/GetBusinessInformation',
			method: 'POST',
			json: {
				ID: this.props.userID,
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
				if(body == "Error"){
					console.log("Error");	
				}else{
					$('#business').val(body.NAME);
					$('#min_age').val(body.MIN_AGE);
					$('#max_age').val(body.MAX_AGE);
					$('#min_people').val(body.MIN_PEOPLE);
					$('#max_people').val(body.MAX_PEOPLE);
					$('#main_category').val(body.MAIN_CATEGORY);
					$('#sub_category').val(body.SUB_CATEGORY);
					$('#address').val(body.ADDRESS);
					$('#township').val(body.TOWNSHIP);
					$('#campus').val(body.CAMPUS);
				}
			}
		}.bind(this));
	},
	updateGeneral: function(){
		request({
			url: 'http://localhost:8080/UpdateBusinessGeneral',
			method: 'POST',
			json: {
				ID: this.props.userID,
				NAME: $('#business').val(),
				MIN_AGE: $('#min_age').val(),
				MAX_AGE: $('#max_age').val(),
				MIN_PEOPLE: $('#min_people').val(),
				MAX_PEOPLE: $('#max_people').val(),
				MAIN_CATEGORY: $('#main_category').val(),
				SUB_CATEGORY: $('#sub_category').val()
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
				if(body != "Success"){
					this.forceUpdate();
				}else{
					console.log(body)		
				}
			}
		}.bind(this));
	},
	updateAddress: function(){
		request({
			url: 'http://localhost:8080/UpdateBusinessAddress',
			method: 'POST',
			json: {
				ID: this.props.userID,
				ADDRESS: $('#address').val(),
				TOWNSHIP: $('#township').val(),
				CAMPUS: $('#campus').val()
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
			}
		}.bind(this));
	},
	updateAll: function(){
		request({
			url: 'http://localhost:8080/UpdateBusinessAll',
			method: 'POST',
			json: {
				ID: this.props.userID,
				NAME: $('#business').val(),
				MIN_AGE: $('#min_age').val(),
				MAX_AGE: $('#max_age').val(),
				MIN_PEOPLE: $('#min_people').val(),
				MAX_PEOPLE: $('#max_people').val(),
				MAIN_CATEGORY: $('#main_category').val(),
				SUB_CATEGORY: $('#sub_category').val(),
				ADDRESS: $('#address').val(),
				TOWNSHIP: $('#township').val()
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
				if(body != "Success"){
					this.forceUpdate();
				}
			}
		}.bind(this));

	},
	render: function(){
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>General Information:</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-6">
						<input type="text" className="form-control" id="business" placeholder="Name"/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Min Age:</span>
							<input type="number" className="form-control" id="min_age"/>
						</div>
					</div>
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Max Age:</span>
							<input type="number" className="form-control" id="max_age"/>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Min People:</span>
							<input type="number" className="form-control" id="min_people"/>
						</div>
					</div>
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Max People:</span>
							<input type="number" className="form-control" id="max_people"/>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Main Category:</span>
							<select className="form-control" id="main_category">
								<Options type="main"/>
							</select>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Sub Category:</span>
							<select className="form-control" id="sub_category">
								<Options typ="sub"/>
							</select>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-5">
						<button type="button"
							className="btn btn-md btn-block btn-primary"
							onClick={this.updateGeneral}>
								Update General Information
						</button>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Address Information:</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-6">
						<input type="text" className="form-control" id="address" placeholder="Address"/>
					</div>
					<div className="col-md-6">
						<input type="text" className="form-control" id="township" placeholder="Township"/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-6">
						<div className="input-group">
							<span className="input-group-addon">Campus</span>
							<select className="form-control" id="campus">
								<option value="New Brunswick">New Brunswick</option>
								<option value="Newark">Newark</option>
								<option value="Camden">Camden</option>
							</select>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-5">
						<button type="button"
							className="btn btn-md btn-block btn-primary"
							onClick={this.updateAddress}>
								Update Address
						</button>
					</div>
					<div className="col-md-offset-2 col-md-5">
						<button type="button"
							className="btn btn-md btn-block btn-success"
							onClick={this.updateAll}>
								Update All
						</button>
					</div>
				</div>	
			</div>
		);
	}
});

module.exports = Settings;
