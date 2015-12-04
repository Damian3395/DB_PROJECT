var React = require('react');

var Coupon = React.createClass({
	getInitialState: function(){
		return({type: "In-Use", 
					business: "Test Business",
					address: "Test Address",
					hours: "Test Hours",
					discount: "Test Discount",
					expires: "Test Expire"});
	},
	componentWillReceiveProps: function(nextProps){
		console.log("Display Coupon Information From Server");
	},
	render: function(){
		return(
			<div className="row">
				<div className="col-md-12">
					<div className="panel panel-success">
						<div className="panel-heading text-center">
							<div><h4><span className="label label-success">{this.state.business}</span></h4></div>
						</div>
						<div className="panel-body">
								<div className="row">
									<div className="col-xs-2">
										<span className="label label-info">Address: {this.state.address}</span>
									</div>
									<div className="col-xs-offset-1 col-xs-2">
										<span className="label label-info">Hours: {this.state.hours}</span>
									</div>
									<div className="col-xs-offset-1 col-xs-2">
										<span className="label label-info">Discount: {this.state.discount}</span>
									</div>
									<div className="col-xs-offset-1 col-xs-2">
										<span className="label label-info">Expires: {this.state.expires}</span>
									</div>
								</div>
						</div>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = Coupon;
