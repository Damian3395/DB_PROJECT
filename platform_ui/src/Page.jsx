var React = require("react");
var Home = require("./Home.jsx");
var About = require("./About.jsx");
var Contact = require("./Contact.jsx");
var UserLogin = require("./UserLogin.jsx");
var BusinessLogin = require("./BusinessLogin.jsx");

var Page = React.createClass({
    render: function(){
        var Body;

        console.log("Prop State %s", this.props.state);
        switch(this.props.state){
            case "Home":
                console.log("Set Body Home Page");
                Body = <Home/>
                break;
            case "About":
                console.log("Set Body About Page");
                Body = <About/>
                break;
            case "Contact":
                console.log("Set Body Contact Page");
                Body = <Contact/>
                break;
            case "User_Login":
                console.log("Set Body User Login Page");
                Body = <UserLogin/>
                break;
            case "Business_Login":
                console.log("Set Body Business Login Page");
                Body = <BusinessLogin/>
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
