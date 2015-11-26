var React = require('react');

var Home = React.createClass({
    setPageState: function(state){
        this.props.stateCallback(state);
    },
    render: function(){
        return(
            <div className="container">
                <div className="row">
                    <div className="col-md-12">
                        <h2 className="text-center"><strong>Welcome To</strong></h2>
                    </div>
                </div>

                <div className="row">
                    <div className="col-md-12">
                        <h1 className="text-center"><strong>RUExploring</strong></h1>
                    </div>
                </div>

                <div className="row">
                    <div className="col-md-6 col-md-offset-3">
                        <h3 className="text-left"><strong>Who Are We?</strong></h3>
                    </div>

                    <div className="col-md-6 col-md-offset-3">
                        <p className="text-justified text-capitalize">
                            We are an organization that brings students and local businesses together by providing retailers
                            the resources and incentives to attract wandering consumers to their product while
                            satisfying a consumers with needs or wants they did not know about.
                            There are more than 67,000 students spread throughout all of Rutger's campuses:
                            New Brunswick, Newark, and Camden. The majority of these students are not locals and lack
                            the knowledge about the thousands of businesses near them. Each year there are roughly 22,812
                            new students that come to Rutgers, providing local businesses a steady flow of consumers.
                            While providing new students the tools they need to socialize and find friends that share
                            the same tastes or interests.
                        </p>
                    </div>
                </div>

                <div className="row">
                    <div className="col-md-offset-1 col-md-4">
                        <h3 className="text-center"><strong>Why Use Our App?</strong></h3>
                    </div>

                    <div className="col-md-offset-2 col-md-4">
                        <h3 className="text-center"><strong>Why Register Your Business?</strong></h3>
                    </div>
                </div>

                <div className="row">
                    <div className="col-md-offset-1 col-md-4">
                        <p className="text-justified text-capitalize">
                            Are you a new student at Rutgers and would like to continue doing your hobbies, but don't know where
                            to go? Having a hard time finding friends to socialize with or finding a hard time finding a place
                            to socialze with your new friends? Would you like to know the best spots to get food close to or
                            near any of Rutger's campuses? Are you a long time student looking for new adventures or activities
                            that are out of the ordinary? Or are you just playing out bored and have nothing better to do? Let
                            our app provide you with random events, places, restraunts, stores, and other fun things to do!
                            Businesses will compete with each other by offering personal discounts and incentives to obtain your
                            business! Allowing you to save money for tuition and enjoy college life!
                        </p>
                    </div>
                    <div className="col-md-offset-2 col-md-4">
                        <p className="text-justified text-capitalize">
                            Are you a new business that is having a hard time finding customers? Would you like to provide
                            potential customers with discounts and other special incentives to bring your compeitors out of business?
                            Going out of business and you need to sell all of your stock products quickly? Get your name out their
                            with our App by allowing us to advertise your name for free, just by registering your business! Don't
                            let your competitors get ahead of you and register now!
                        </p>
                    </div>
                </div>

                <div className="row">
                    <div className="col-md-offset-2 col-md-2">
                        <button type="button" className="btn btn-success btn-lg btn-block" onClick={this.setPageState.bind(this, "Register_User")}>Register Now</button>
                    </div>

                    <div className="col-md-offset-4 col-md-2">
                        <button type="button" className="btn btn-success btn-lg btn-block" onClick={this.setPageState.bind(this, "Register_Business")}>Register Now</button>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = Home;
