var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var request = require('request');

var Settings = React.createClass({
	getInitialState: function(){
		return({student: true, ruid: "", degree: "", campus: "", year: ""});
	},
	componentWillMount: function(){
		request({
			url: 'http://localhost:8080/GetUserInformation',
			method: 'POST',
			json: {
				ID: this.props.userID
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(body)
				$('#firstName').val(body.FIRST_NAME);
				$('#lastName').val(body.LAST_NAME);
				$('#age').val(body.AGE);
				$('#gender').val(body.GENDER);
				$('#address').val(body.ADDRESS);
				$('#township').val(body.TOWNSHIP);
				$('#state').val(body.STATE);
				$('#zipcode').val(body.ZIPCODE);
			}
		}.bind(this));
		request({
			url: 'http://localhost:8080/GetStudentInformation',
			method: 'POST',
			json: {
				ID: this.props.userID
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				if(body !== "Error: User Not Found"){
					$('#ruid').val(body.RUID);
					$('#campus').val(body.CAMPUS);
					$('#degree').val(body.DEGREE);
					$('#year').val(body.YEAR);
					this.setState({student: true, ruid: body.RUID, campus: body.CAMPUS, degree: body.DEGREE, year: body.YEAR});
				}else{
					this.setState({student: false});
				}
			}
		}.bind(this));
	},
	updateGeneral: function(){
		request({
			url: 'http://localhost:8080/UpdateUserGeneral',
			method: 'POST',
			json: {
				ID: this.props.userID,
				FIRST_NAME: $('#firstName').val(),
				LAST_NAME: $('#lastName').val(),
				AGE: $('#age').val(),
				GENDER: $('#gender').val()
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
	updateAddress: function(){
		request({
			url: 'http://localhost:8080/UpdateUserAddress',
			method: 'POST',
			json: {
				ID: this.props.userID,
				ADDRESS: $('#address').val(),
				TOWNSHIP: $('#township').val(),
				STATE: $('#state').val(),
				ZIPCODE: $('#zipcode').val()
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
	updateStudent: function(){
		request({
			url: 'http://localhost:8080/UpdateUserStudent',
			method: 'POST',
			json: {
				ID: this.props.userID,
				RUID: $('#ruid').val(),
				DEGREE: $('#degree').val(),
				CAMPUS: $('#campus').val(),
				YEAR: $('#year').val()
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
	updateAll: function(){
		var ruid, degree, campus, year;
		if(!this.state.student){
			ruid = "remove";
			degree = "";
			campus = "";
			year = "";
		}else{
			ruid = $('#ruid').val();
			degree = $('#degree').val();
			campus = $('#campus').val();
			year = $('#year').val();
		}
		request({
			url: 'http://localhost:8080/UpdateUserAll',
			method: 'POST',
			json: {
				ID: this.props.userID,
				FIRST_NAME: $('#firstName').val(),
				LAST_NAME: $('#lastName').val(),
				AGE: $('#age').val(),
				GENDER: $('#gender').val(),
				ADDRESS: $('#address').val(),
				TOWNSHIP: $('#township').val(),
				STATE: $('#state').val(),
				ZIPCODE: $('#zipcode').val(),
				RUID: ruid,
				DEGREE: degree,
				CAMPUS: campus,
				YEAR: year
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
	studentChecked: function(){
		if(!this.state.student){
			$('#ruid').val(this.state.ruid);
			$('#campus').val(this.state.campus);
			$('#degree').val(this.state.degree);
			$('#year').val(this.state.year);
		}
		this.setState({student: !this.state.student});
	},
	render: function(){
		var Student;
		if(this.state.student){
			Student =
				<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Student Information:</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-6">
						<input type="text" className="form-control" id="ruid"/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-7">
						<div className="input-group">
							<span className="input-group-addon">Select Your Degree</span>
							<select className="form-control" id="degree">
								<option value="Computer Science">Computer Science</option>	
								<option value="Psychology">Psychology</option>	
								<option value="Engineering">Engineering</option>	
								<option value="Mathematics">Mathematics</option>	
							</select>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-7">
						<div className="input-group">
							<span className="input-group-addon">Select Your Campus</span>
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
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Year</span>
							<select className="form-control" id="year">
								<option value="Freshmen">Freshmen</option>	
								<option value="Sophmore">Sophmore</option>	
								<option value="Junior">Junior</option>	
								<option value="Senior">Senior</option>	
							</select>
						</div>
					</div>
				</div>
				<br/>
				<div classname="row">
					<div className="col-md-6">
						<button type="button"
							className="btn btn-primary btn-md"
							onClick={this.updateStudent}>
								Update Student Information
						</button>
					</div>
				</div>
			</div>
			;
		}

		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>General Information:</strong></h3>	
					</div>
				</div>
				<div className="row">
					<div className="col-md-4">
						<input type="text" className="form-control" id="firstName"/>
					</div>
					<div className="col-md-4">
						<input type="text" className="form-control" id="lastName"/>	
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Gender</span>
							<select className="form-control" id="gender">
								<option value="Male">Male</option>	
								<option value="Female">Female</option>	
							</select>
						</div>
					</div>
					<div className="col-md-4">
						<input type="number" className="form-control" id="age"/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-6">
						<button type="button"
							className="btn btn-md btn-primary"
							onClick={this.updateGeneral}>
								Update General Information
						</button>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Address:</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-offset-md-2 col-md-6">
						<input type="text" className="form-control" id="address"/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<input type="text" className="form-control" id="township"/>
					</div>
					<div className="col-md-4">
						<input type="text" className="form-control" id="state"/>
					</div>
					<div className="col-md-4">
						<input type="number" className="form-control" id="zipcode"/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-6">
						<button type="button"
							className="btn btn-md btn-primary"
							onClick={this.updateAddress}>
								Update Address
							</button>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-6">
						<div className="checkbox">
							<label><input type="checkbox" 
										id="isStudent" 
										onClick={this.studentChecked} 
										checked={this.state.student}/>
											RU a Rutgers Student?
							</label>
						</div>
					</div>
				</div>
				{Student}
				<br/>
				<div className="row">
					<div className="col-md-6">
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
