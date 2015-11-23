var React = require('react');

var NavBar = React.createClass({
    setPageState: function(state){
        console.log("Setting State %s", state);
        this.props.stateCallback(state);
    },
    render: function(){
        return(
            <div>
                <nav className="navbar navbar-default">
                    <div className="container-fluid">
                        <div className="navbar-header"></div>
                        <div></div>
                        <div className="navbar-collapse collapse">
                            <form className="navbar-form navbar-left">
                                <button className="btn btn-default" onClick={this.setPageState.bind(this, "Home")}>RUExploring</button>
                                <button className="btn btn-default" onClick={this.setPageState.bind(this, "About")}>About</button>
                                <button className="btn btn-default" onClick={this.setPageState.bind(this, "Contact")}>Contact</button>
                            </form>
                            <form className="navbar-form navbar-right">
                                <button className="btn btn-default" onClick={this.setPageState.bind(this, "User_Login")}>User Login</button>
                                <button className="btn btn-default" onClick={this.setPageState.bind(this, "Business_Login")}>Business Login</button>
                            </form>
                        </div>
                    </div>
                </nav>
            </div>
        );
    }
});

module.exports = NavBar;
