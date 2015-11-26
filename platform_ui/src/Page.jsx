var React = require("react");
var Home = require("./pages/misc/Home.jsx");
var About = require("./pages/misc/About.jsx");
var Contact = require("./pages/misc/Contact.jsx");
var UserLogin = require("./pages/login/UserLogin.jsx");
var RegisterUser = require("./pages/register/RegisterUser.jsx");
var BusinessLogin = require("./pages/login/BusinessLogin.jsx");
var RegisterBusiness = require("./pages/register/RegisterBusiness.jsx");
var App = require("./pages/main/App.jsx");

var Page = React.createClass({
	setPageState: function(state){
		this.props.stateCallback(state);
	},
    render: function(){
        var Body;

        switch(this.props.state){
            case "Home":
                Body = <Home stateCallback={this.setPageState}/>
                break;
            case "About":
                Body = <About/>
                break;
            case "Contact":
                Body = <Contact/>
                break;
            case "User_Login":
                Body = <UserLogin stateCallback={this.setPageState}/>
                break;
            case "Business_Login":
                Body = <BusinessLogin stateCallback={this.setPageState}/>
                break;
	    case "Register_User":
		Body = <RegisterUser stateCallback={this.setPageState}/>
		break;
            case "Register_Business":
		Body = <RegisterBusiness stateCallback={this.setPageState}/>
		break;
            case "App":
		Body = <App stateCallback={this.setPageState}/>
		break;
            default:
                Body = <Home/>
                console.log("Error Invalid Page State");
        }

        return(
            <div>
                {Body}
            </div>
        );
    }
});

module.exports = Page;
