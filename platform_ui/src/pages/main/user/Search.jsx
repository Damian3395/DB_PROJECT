var React = require('react');
var $ = require('jquery');
var request = require('request');

var Search = React.createClass({
	getInitialState: function(){
		return({query: "", main_category: "", sub_category: "", min_age: "", max_age: "", min_people: "", max_people: "", campus: ""});
	},
	getRandom: function(){
		console.log("Random Search Result");
	},
	getSearchResult: function(){
		console.log("Getting Search Result");
	},
	render: function(){
		return(
			<div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Search For Coupons</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-12">
						<div className="input-group">
							<span className="input-group-addon">Search For Key Words:</span>
							<input type="text" className="form-control" id="query"/>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Main Category: </span>
							<select	className="form-control" id="main_category">
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
							<span className="input-group-addon">Sub Category: </span>
							<select	className="form-control" id="sub_category">
								<option value="Four">Four</option>
								<option value="Five">Five</option>
								<option value="Six">Six</option>
							</select>
						</div>	
					</div>
				</div>	
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Campus: </span>
							<select	className="form-control" id="campus">
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
							<span className="input-group-addon">Min Age:</span>
							<input type="number" className="form-control" id="min_age"/>
						</div>	
					</div>
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Max Age:</span>
							<input type="number" className="form-control" id="max_age"/>
						</div>	
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Min People:</span>
							<input type="number" className="form-control" id="min_people"/>
						</div>	
					</div>
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Max People:</span>
							<input type="number" className="form-control" id="max_people"/>
						</div>	
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<button type="button"
							className="btn btn-lg btn-block btn-success"
							onClick={this.getRandom}>
								Random
						</button>
					</div>
					<div className="col-md-offset-4 col-md-4">
						<button type="button"
							className="btn btn-lg btn-block btn-primary"
							onClick={this.getSearchResult}>
								Search
						</button>
					</div>	
				</div>
				<br/>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Coupons Found</strong></h3>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = Search;
