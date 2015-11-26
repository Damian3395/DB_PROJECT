var React = require('react');

var Footer = React.createClass({
	render: function(){
		return(
			<div className="container">
				<div className="row">
					<div className="col-md-12">
						<p className="text-right">
							Damian Debkowski & Nicole Yson -
							Rutgers School of Arts and Sciences -
							Principles of Information & Data Management -
							2015
						</p>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = Footer;
