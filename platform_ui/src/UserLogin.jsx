var React = require('react');

var UserLogin = React.createClass({
	loginUser: function(state){
		console.log("Login User");
		this.props.stateCallback(state);
	},
	registerNewUser: function(state){
		this.props.stateCallback(state);
	},
    render: function(){
        return(
            <div className="container">
				<div className="row">
					<div className="col-md-4 col-md-offset-4 text-center">
						<div className="panel panel-default">
							<div className="panel-heading">
								<h2 className="panel-title text-center"><strong>User Login</strong></h2>
							</div>
							<div className="panel-body">
								<div className="row">
									<div className="col-xs-offset-1 col-xs-10">
									<input className="form-control" type="text" id="userID" placeholder="username"/>
									</div>
								</div>
								<br/>
								<div className="row">
									<div className="col-xs-offset-1 col-xs-10">
										<input className="form-control" type="password" id="password"/>
									</div>
								</div>
								<br/>
								<div className="btn-group btn-group-justified">
									<div className="btn-group">
										<button className="btn btn-success btn-lg" 
											type="button" 
											onClick={this.loginUser.bind(this, "App")}>
												Login
										</button>
									</div>
									<div className="btn-group">		
										<button className="btn btn-primary btn-lg" 
											type="button" 
											onClick={this.registerNewUser.bind(this, "Register_User")}>
												Register
											</button>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
        );
    }
});

module.exports = UserLogin;
