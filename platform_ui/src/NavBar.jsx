var React = require('react');

var NavBar = React.createClass({
    setPageState: function(state){
        this.props.stateCallback(state, this.props.type, this.props.userID);
    },
    render: function(){
		var login_App =
			<div className="col-md-offset-4 col-md-3">
				<button className="btn btn-default navbar-btn"
					onClick={this.setPageState.bind(this, "User_Login")}>
						User Login
				</button>
				<button className="btn btn-default navbar-btn" 
					onClick={this.setPageState.bind(this, "Business_Login")}>
						Business Login
				</button>		
			</div> 
			;
		if(this.props.userID !== ""){
			login_App =
				<div className="col-md-offset-4 col-md-3">
					<button className="btn btn-default navbar-btn"
						onClick={this.setPageState.bind(this, "App")}>
							Application
					</button>
				</div>
			;
		}
        return(
            <div>
                <nav className="navbar navbar-default">
                    <div className="container-fluid">
						<div className="row">
							<div className="col-md-offset-1 col-md-4">
		<button className="btn btn-default navbar-btn" onClick={this.setPageState.bind(this, "Home")}>RUExploring</button>
        <button className="btn btn-default nvabar-btn" onClick={this.setPageState.bind(this, "About")}>About</button>
        <button className="btn btn-default navbar-btn" onClick={this.setPageState.bind(this, "Contact")}>Contact</button>	
							</div>
							<div>
								{login_App}
							</div>
						</div>
                    </div>
                </nav>
            </div>
        );
    }
});

module.exports = NavBar;
