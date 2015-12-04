var React = require('react');
var $ = require('jquery');
var request = require('request');

var Settings = React.createClass({
	getInitialState: function(){
		return({
			first_name: "Test Name",
			last_name: "Test Last",
			age: "Test Age",
			address: "Test Address",
			township: "Test Township",
			state: "Test State",
			zipcode: "Test Zipcode",
			student: true,
			ruid: "Test RUID",
			degree: "Test Degree",
			campus: "Test Campus",
			year: "Test Year"
		});
	},
	componentWillMount: function(){
		console.log("Get Information From Server");
	},
	updateGeneral: function(){
		console.log("Update General Information");
	},
	updateAddress: function(){
		console.log("Updating Address Information");
	},
	updateStudent: function(){
		console.log("Updating Student Information");
	},
	updateAll: function(){
		console.log("Updating All User Information");
	},
	studentChecked: function(){
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
						<input type="text" className="form-control" id="ruid" placeholder={this.state.ruid}/>
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
						<input type="text" className="form-control" id="firstName" placeholder={this.state.first_name}/>
					</div>
					<div className="col-md-4">
						<input type="text" className="form-control" id="lastName" placeholder={this.state.last_name}/>	
					</div>
					<div className="col-md-4">
						<input type="number" className="form-control" id="age" placeholder={this.state.age}/>
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
						<input type="text" className="form-control" id="address" placeholder={this.state.address}/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<input type="text" className="form-control" id="township" placeholder={this.state.township}/>
					</div>
					<div className="col-md-4">
						<input type="text" className="form-control" id="state" placeholder={this.state.state}/>
					</div>
					<div className="col-md-4">
						<input type="number" className="form-control" id="zipcode" placeholder={this.state.zipcode}/>
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
							<label><input type="checkbox" id="isStudent" onClick={this.studentChecked} checked={this.state.student}/>RU a Rutgers Student?</label>
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
