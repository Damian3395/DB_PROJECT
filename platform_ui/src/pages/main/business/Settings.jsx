var React = require('react');
var $ = require('jquery');
var request = require('request');

var Settings = React.createClass({
	getInitialState: function(){
		return({
			business: "Test Business",
			address: "Test Address",
			township: "Test township",
			min_age: "Test Min Age",
			max_age: "Test Max Age",
			min_people: "Test Min People",
			max_people: "Test Max People",
			main_category: "Test Main Category",
			sub_category: "Test Sub Category"
		});
	},
	componentWillMount: function(){
		console.log("Get Information From Server");
	},
	updateGeneral: function(){
		console.log("Update General Information");
	},
	updateAddress: function(){
		console.log("Update Address Information");
	},
	updateAll: function(){
		console.log("Update All Information");
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
						<input type="text" className="form-control" id="business" placeholder={this.state.business}/>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Min Age:</span>
							<input type="number" className="form-control" id="min_age" placeholder={this.state.min_age}/>
						</div>
					</div>
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Max Age:</span>
							<input type="number" className="form-control" id="max_age" placeholder={this.state.max_age}/>
						</div>
					</div>
				</div>
				<br/>
<div className="row">
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Min People:</span>
							<input type="number" className="form-control" id="min_people" placeholder={this.state.min_people}/>
						</div>
					</div>
					<div className="col-md-5">
						<div className="input-group">
							<span className="input-group-addon">Max People:</span>
							<input type="number" className="form-control" id="max_people" placeholder={this.state.max_people}/>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Main Category:</span>
							<select className="form-control" id="main_category">
								<option value="One">One</option>
								<option value="Two">Two</option>
								<option value="Three">Three</option>
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
								<option value="Four">Four</option>
								<option value="Five">Five</option>
								<option value="Six">Six</option>
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
						<input type="text" className="form-control" id="address" placeholder={this.state.address}/>
					</div>
					<div className="col-md-6">
						<input type="text" className="form-control" id="township" placeholder={this.state.township}/>
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
