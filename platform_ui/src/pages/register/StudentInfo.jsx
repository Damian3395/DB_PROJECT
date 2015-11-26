var React = require('react');
var $ = require('jquery');

var StudentInfo = React.createClass({
	render: function(){
		if(this.props.isStudent){
			return (
				<div>
					<div className="row">
						<div className="col-xs-offset-3 col-xs-6">
							<input className="form-control" type="text" id="ruid" placeholder="RUID"/>
						</div>
					</div>
					<br/>
					<div className="row">
						<div className="col-xs-offset-2 col-xs-8">
							<div className="input-group">
								<span className="input-group-addon">Select Your Degree:</span>
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
						<div className="col-xs-offset-2 col-xs-8">
							<div className="input-group">
								<span className="input-group-addon">Select Your Campus:</span>
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
						<div className="col-xs-offset-2 col-xs-8">
							<div className="input-group">
								<span className="input-group-addon">Select Your Year:</span>
								<select className="form-control" id="year">
									<option value="Freshmen">Freshmen</option>
									<option value="Sophmore">Sophmore</option>
									<option value="Junior">Junior</option>
									<option value="Senior">Senior</option>
								</select>
							</div>
						</div>
					</div>
				</div>
			);
		}else{
			return false;
		}
	}
});

module.exports = StudentInfo;
