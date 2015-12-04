var React = require('react');

var Manage = React.createClass({
	getInitialState: function(){
		return{student: false, createType: "", day: "", month: "", year: ""};
	},
	isStudent: function(){
		this.setState({student: !this.state.student});
	},
	render: function(){
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
								<option value="one">One</option>
								<option value="two">Two</option>
								<option value="three">Three</option>
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
				</div>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Current Coupons</strong></h3>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = Manage;
