var React = require('react');

var Footer = React.createClass({
	render: function(){
		return(
			<div className="container">
				<div className="row">
					<div className="col-md-12">
						<footer className="text-right">
							<p>
								Damian Debkowski & Nicole Yson -
								Rutgers School of Arts and Sciences -
								Principles of Information & Data Management -
								2015
							</p>
						</footer>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = Footer;
