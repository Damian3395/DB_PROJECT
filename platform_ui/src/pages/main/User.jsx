var React = require('react');
var Profile = require('./user/Profile.jsx');
var Search = require('./user/Search.jsx');
var Settings = require('./user/Settings.jsx');

var User = React.createClass({
	getInitialState: function(){
		return {Pane: "Profile"};
	},
	setPane: function(paneState){
		this.setState({Pane: paneState});
	},
	logoutUser: function(){
		this.props.stateCallback("Home", "", "");
	},
	render: function(){
		var Pane;
		switch(this.state.Pane){
			case "Search":
				Pane = <Search userID={this.props.userID}/>
				break;
			case "Profile":
				Pane = <Profile userID={this.props.userID}/>
				break;
			case "Settings":
				Pane = <Settings userID={this.props.userID}/>
				break;
			default:
				Pane = <Profile userID={this.props.userID}/>
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
								onClick={this.setPane.bind(this, "Search")}>
									Search
							</button>
							<br/>
							<button type="button"
								className="btn btn-lg btn-primary btn-block"
								onClick={this.setPane.bind(this, "Profile")}>
									Profile
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
								onClick={this.logoutUser}>
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

module.exports = User;
