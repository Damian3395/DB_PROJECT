var React = require('react');

var Footer = React.createClass({
	render: function(){
		return(
			<div>
				<footer className="footer">
					<div className="container">
						<div className="row">
							<p className="text-muted">
								Damian Debkowski & Nicole Yson -
								Rutgers School of Arts and Sciences -
								Principles of Information & Data Management -
								2015
							</p>
						</div>
					</div>
				</footer>
			</div>
		);
	}
});

module.exports = Footer;
