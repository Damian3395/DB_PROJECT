var React = require('react');
var StudentInfo = require('./StudentInfo.jsx');

var RegisterUser = React.createClass({
	getInitialState: function(){
		return({student: false});
	},
	register: function(){
		console.log("Register New User");
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
											<input type="checkbox" id="Student" onClick={this.studentChecked}/>
										</div>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-12">
										<input className="form-control" type="text" id="Address" placeholder="Street Address"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-12">
										<input className="form-control" type="text" id="Township" placeholder="Township"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-5">
										<input className="form-control" type="text" id="State" placeholder="State"/>
									</div>
									<div className="col-xs-offset-1 col-xs-5">
										<div className="input-group">
											<span className="input-group-addon">Zipcode</span>
											<input className="form-control" type="number" id="ZipCode"/>
										</div>
									</div>
								</div>
								<br/>
								<StudentInfo isStudent={this.state.student}>
								<div>
									<div className="row">
										<div className="col-xs-offset-3 col-xs-6">
											<input className="form-control" type="text" id="RUID" placeholder="RUID"/>
										</div>
									</div>
									<br/>
									<div className="row">
										<div className="col-xs-offset-3 col-xs-6">
											<div className="btn-group">
											<div className="btn-group">
												<button type="button" className="btn btn-primary dropdown-toggle" id="Degree"
													data-toggle="dropdown" 
													aria-haspopup="true" 
													aria-expanded="false">
														Degree
														<span className="caret"></span>
												</button>
												<ul className="dropdown-menu">
													<li><a href="#">Computer Science</a></li>
													<li><a href="#">Psychology</a></li>
													<li><a href="#">Engineering</a></li>
													<li><a href="#">Mathematics</a></li>
												</ul>
											</div>

											<div className="btn-group">
												<button type="button" className="btn btn-primary dropdown-toggle" id="Campus"
													data-toggle="dropdown"
													aria-haspopup="true"
													aria-expanded="false">
														Campus
														<span className="caret"></span>
												</button>
												<ul className="dropdown-menu">
													<li><a href="#">New Brunswick</a></li>
													<li><a href="#">Newark</a></li>
													<li><a href="#">Camden</a></li>
												</ul>
											</div>

											<div className="btn-group">
												<button type="button" className="btn btn-primary dropdown-toggle" id="Year"
													data-toggle="dropdown"
													aria-haspopup="true"
													aria-expanded="false">
														Year
														<span className="caret"></span>
												</button>
												<ul className="dropdown-menu">
													<li><a href="#">Freshmen</a></li>
													<li><a href="#">Sophmore</a></li>
													<li><a href="#">Junior</a></li>
													<li><a href="#">Senior</a></li>
												</ul>
											</div>
											</div>
										</div>
										<br/>
										<div className="row">
										
										</div>
									</div>
								</div>
								</StudentInfo>
								<div className="row">
									<div className="col-xs-offset-8 col-xs-4">
										<button className="btn btn-success btn-lg btn-block" 
											type="button" onClick={this.register}>
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
