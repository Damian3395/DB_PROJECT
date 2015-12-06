var React = require('react');
var $ = require('jquery');
var request = require('request');
var Error = require('../misc/Error.jsx');

var RegisterBusiness = React.createClass({
	getInitialState: function(){
		return({error: ""});
	},
	registerBusiness: function(state){
		var userID = $('#userID').val();
		var passwordOne = $('#password_One').val();
		var passwordTwo = $('#password_Two').val();
		var name = $('#name').val();
		var address = $('#address').val();
		var township = $('#township').val();
		var campus = $('#campus').val();
		var min_age = $('#min_Age').val();
		var max_age = $('#max_Age').val();
		var min_people = $('#min_People').val();
		var max_people = $('#max_People').val();
		var main_category = $('#main_Category').val();
		var sub_category = $('#sub_Category').val();	

		if(passwordOne !== passwordTwo){
			this.setState({error: "Password Field 1 and 2 Do Not Match!"});
		}else{
			request({
				url: 'http://localhost:8080/RegisterBusiness',
				method: 'POST',
				json: {
					NAME: name,
					ADDRESS: address,
					TOWNSHIP: township,
					CAMPUS: campus,
					MIN_AGE: min_age,
					MAX_AGE: max_age,
					MIN_PEOPLE: min_people,
					MAX_PEOPLE: max_people,
					MAIN_CATEGORY: main_category,
					SUB_CATEGORY: sub_category,
					USER_ID: userID,
					PASSWORD: passwordOne
				}
			}, function(error, response, body){
				if(error){
					console.log(error);
				}else{
					console.log(response.statusCode, body);
					if(body == "Success"){
						this.props.stateCallback("Business_Login");
					}else{
						console.log(body);
					}

				}
			}.bind(this));
		}
	},
	render: function(){
		return(
			<div className="container">
				<div className="row">
					<div className="col-md-offset-3 col-md-6">
						<div className="panel panel-default">
							<div className="panel-heading">
								<h1 className="panel-title text-center"><strong>Registering Business</strong></h1>
							</div>
							<div className="panel-body">
								<Error error={this.state.error}/>
								<br/>
								<div className="row">
									<div className="col-xs-offset-3 col-xs-6">
										<input className="form-control" type="text" id="name" placeholder="Business Name"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-12">
										<input className="form-control" type="text" id="address" placeholder="Address"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-8">
										<div className="input-group">
											<span className="input-group-addon">Campus:</span>
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
									<div className="col-xs-8">
										<div className="input-group">
											<span className="input-group-addon">Township:</span>
											<input type="text" className="form-control" id="township"/>
										</div>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-5">
										<div className="input-group">
											<span className="input-group-addon">Min Age:</span>
											<input className="form-control" type="number" id="min_Age"/>
										</div>
									</div>
									<div className="col-xs-offset-1 col-xs-5">
										<div className="input-group">
											<span className="input-group-addon">Max Age:</span>
											<input className="form-control" type="number" id="max_Age"/>
										</div>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-5">
										<div className="input-group">
											<span className="input-group-addon">Min People:</span>
											<input className="form-control" type="number" id="min_People"/>
										</div>
									</div>
									<div className="col-xs-offset-1 col-xs-5">
										<div className="input-group">
											<span className="input-group-addon">Max People:</span>
											<input className="form-control" type="number" id="max_People"/>
										</div>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-8">
										<div className="input-group">
											<span className="input-group-addon">Main Category:</span>
											<select className="form-control" id="main_Category">
												<option value="Food">Food</option>
												<option value="Entertainment">Entertainment</option>
												<option value="Sport">Sport</option>
												<option value="Religious">Religious</option>
												<option value="Store">Store</option>
												<option value="Other">Other</option>
											</select>
										</div>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-8">
										<div className="input-group">
											<span className="input-group-addon">Sub Category:</span>
											<select className="form-control" id="sub_Category">
												<option value="1">1</option>
												<option value="2">2</option>
												<option value="3">3</option>
												<option value="4">4</option>
												<option value="5">5</option>
												<option value="6">6</option>
											</select>
										</div>
									</div>
								</div>
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
											type="button"
											onClick={this.registerBusiness.bind(this, "Business_Login")}>
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

module.exports = RegisterBusiness;
