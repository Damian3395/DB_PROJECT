var React = require('react');

var Contact = React.createClass({
    render: function(){
        return(
            <div className="container">
				<div className="row">
					<div className="col-md-12">
						<h1 className="text-left"><strong>Contact Information:</strong></h1>
					</div>
				</div>

				<div className="row">
					<div className="col-md-4">
						<p className="text-left"><strong>PLEASE DON'T CONTACT US</strong></p>
						<p className="text-left">Email: nicole.yson@rutgers.edu</p>
						<p className="text-left">Email: damian.debkowski@rutgers.edu</p>
						<p className="text-left">Email: imielins@cs.rutgers.edu</p>
						<p className="text-left">Email: undergrad-office@cs.rutgers.edu</p>
					</div>
				</div>
            </div>
        )
    }
});

module.exports = Contact;
