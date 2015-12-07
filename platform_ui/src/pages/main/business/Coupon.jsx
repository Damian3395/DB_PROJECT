var React = require('react');

var Coupon = React.createClass({
	render: function(){
		return(
			<div className="row">
				<div className="col-md-12">
					<div className="panel panel-success">
						<div className="panel-heading text-center">
							<div><h4><span className="label label-success">Coupon {this.props.name}</span></h4></div>
						</div>
						<div className="panel-body">
								<div className="row">
									<div className="col-xs-8">
										<span className="label label-info">Discount: {this.props.type}</span>
									</div>
									<div className="col-xs-4">
										<span className="label label-info">
											Expires: {this.props.month}-{this.props.day}-{this.props.year}
										</span>
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
