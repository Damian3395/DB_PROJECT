var React = require('react');
var Manage = require('./business/Manage.jsx');
var Analytics = require('./business/Analytics.jsx');
var Settings = require('./business/Settings.jsx');

var Business = React.createClass({
	getInitialState: function(){
		return {Pane: "Manage"};
	},
	setPane: function(paneState){
		this.setState({Pane: paneState});
	},
	logoutBusiness: function(){
		this.props.stateCallback("Home", "", "");
	},
	render: function(){
		var Pane;
		switch(this.state.Pane){
			case "Manage":
				Pane = <Manage userID={this.props.userID}/>
				break;
			case "Analytics":
				Pane = <Analytics userID={this.props.userID}/>
				break;
			case "Settings":
				Pane = <Settings userID={this.props.userID}/>
				break;
			default:
				Pane = <Manage userID={this.props.userID}/>
		}

		return(
			<div className="row">
				<div className="col-offset-lg-1 col-lg-4">
					<div className="panel panel-primary">
						<div className="panel-heading">
							<h1 className="text-center"><strong>Menu</strong></h1>
						</div>
						<div className="panel-body">
							<button type="button"
								className="btn btn-lg btn-primary btn-block"
								onClick={this.setPane.bind(this, "Manage")}>
									Manage
							</button>	
							<br/>
							<button type="button"
								className="btn btn-lg btn-primary btn-block"
								onClick={this.setPane.bind(this, "Analytics")}>
									Analytics
							</button>
							<br/>
							<button type="button"
								className="btn btn-lg btn-primary btn-block"
								onClick={this.setPane.bind(this, "Settings")}>
									Settings
							</button>
							<br/>
							<button type="button"
								className="btn btn-lg btn-primary btn-block"
								onClick={this.logoutBusiness}>
									Logout
							</button>
						</div>
					</div>
				</div>
				<div className="col-offset-lg-1 col-lg-7">
					<div className="panel panel-default">
						<div className="panel-heading">
							<h1 className="text-center"><strong>{this.state.Pane}</strong></h1>
						</div>
						<div className="panel-body">
							{Pane}
						</div>
					</div>		
				</div>
			</div>
		);
	}
});

module.exports = Business;
