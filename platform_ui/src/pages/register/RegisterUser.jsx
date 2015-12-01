var React = require('react');
var $ = require('jquery')
var request = require('request')
var StudentInfo = require('./StudentInfo.jsx');
var Error = require('../misc/Error.jsx');

var RegisterUser = React.createClass({
	getInitialState: function(){
		return({student: false, error: ""});
	},
	registerUser: function(state){
		var userID = $('#userID').val();
		var passwordOne = $('#password_One').val();
		var passwordTwo = $('#password_Two').val();
		var firstName = $('#firstName').val();
		var lastName = $('#lastName').val();
		var age = $('#age').val();
		var address = $('#address').val();
		var township = $('#township').val();
		var state = $('#state').val();
		var zipcode = $('#zipcode').val();
		var student;
		var ruid;
		var degree;
		var campus;
		var year;

		if(this.state.student){
			student = "true";
			ruid = $('#ruid').val();
			degree = $('#degree').val();
			campus = $('#campus').val();
			year = $('#year').val();
		}else{
			student = "false";
			ruid = "null";
			degree = "null";
			campus = "null";
			year = "null";
		}
		
		if(passwordOne !== passwordTwo){
			this.setState({error: "Password Field 1 and 2 Do Not Match!"});
		}else{
			request({
				url: 'http://localhost:8080/RegisterUser',
				method: 'POST',
				json: {
					FIRST_NAME: firstName,
					LAST_NAME: lastName,
					AGE: age,
					STUDENT: student,
					ADDRESS: address,
					TOWNSHIP: township,
					STATE: state,
					RUID: ruid,
					ZIPCODE: zipcode,
					DEGREE: degree,
					CAMPUS: campus, 
					YEAR: year,
					USER_ID: userID,
					PASSWORD: passwordOne
				}		
			}, function(error, response, body){
				if(error){
					console.log(error);
				}else{
					console.log(response.statusCode, body);
					console.log("Creating cookie");
					this.props.stateCallback("App");
				}
			}.bind(this));
		}	
	},
	studentChecked: function(){
		this.setState({student: !this.state.student});
	},
	render: function(){
		return(
			<div className="container">
				<div className="row">
					<div className="col-md-offset-3 col-md-6">
						<div className="panel panel-default">
							<div className="panel-heading">
								<h1 className="panel-title text-center"><strong>Registering User</strong></h1>
							</div>
							<div className="panel-body">
								<Error error={this.state.error}/>
								<br/>
								<div className="row">
									<div className="col-xs-5">
										<input className="form-control" type="text" id="firstName" placeholder="First Name"/>
									</div>
									<div className="col-xs-offset-1 col-xs-5">	
										<input className="form-control" type="text" id="lastName" placeholder="Last Name"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-4">
										<div className="input-group">
											<span className="input-group-addon">Age:</span>
											<input className="form-control" type="number" id="age"/>
										</div>	
									</div>
									<div className="col-xs-offset-2 col-xs-3">
										<div className="input-group">
											<span className="input-group-addon">Are You A Rutgers Student?</span>
										</div>
									</div>
									<div className="col-xs-offset-2 col-xs-1">
										<input type="checkbox" id="student" onClick={this.studentChecked}/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-12">
										<input className="form-control" type="text" id="address" placeholder="Street Address"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-12">
										<input className="form-control" type="text" id="township" placeholder="Township"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-5">
										<input className="form-control" type="text" id="state" placeholder="State"/>
									</div>
									<div className="col-xs-offset-1 col-xs-5">
										<div className="input-group">
											<span className="input-group-addon">Zipcode</span>
											<input className="form-control" type="number" id="zipcode"/>
										</div>
									</div>
								</div>
								<br/>
								<StudentInfo isStudent={this.state.student}/>
								<br/>
								<div className="row">
									<div className="col-xs-offset-2 col-xs-8">
										<input className="form-control" type="text" id="userID" placeholder="User Name"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-offset-2 col-xs-8">
										<label>Password:</label>
										<input className="form-control" type="password" id="password_One"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-offset-2 col-xs-8">
										<label>Re-Enter Password:</label>
										<input className="form-control" type="password" id="password_Two"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-offset-8 col-xs-4">
										<button className="btn btn-success btn-lg btn-block" 
											type="button" onClick={this.registerUser.bind(this, "App")}>
												Register
										</button>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		);	
	}
});

module.exports = RegisterUser;
