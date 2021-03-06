var React = require('react');
var $ = require('jquery');
var request = require('request');

var BusinessLogin = React.createClass({
	loginBusiness: function(state){
		request({
			url: 'http://www.ruexploring.com:80/LoginBusiness',
			method: 'POST',
			json: {
				ID: $('#userID').val(),
				PASSWORD: $('#password').val()
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				if(body == "Success"){
					this.props.stateCallback("App", "Business", $('#userID').val());
				}else{
					console.log("Display Error");
				}
			}
		}.bind(this));		
	},
	registerNewBusiness: function(state){
		this.props.stateCallback(state, "Register_Business", "");	
	},
	render: function(){
        return(
            <div className="container">
				<div className="row">
					<div className="col-md-4 col-md-offset-4 text-center">
                        <div className="panel panel-default">
                            <div className="panel-heading">
                                <h2 className="panel-title text-center"><strong>Business Login</strong></h2>
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
											onClick={this.loginBusiness.bind(this, "App")}>
												Login
										</button>
                                    </div>
                                    <div className="btn-group">
										<button className="btn btn-primary btn-lg" 
											type="button" 
											onClick={this.registerNewBusiness.bind(this, "Register_Business")}>
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

module.exports = BusinessLogin;
