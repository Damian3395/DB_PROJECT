var React = require('react');
var Coupon = require('./Coupon.jsx');

var Display = React.createClass({
	getInitialState: function(){
		return{tickets: []};	
	},
	componentWillReceiveProps: function(nextProps){
		if(nextProps.tickets.length != this.props.tickets.length){
			console.log("update display state")
			this.setState({tickets: nextProps.tickets});
		}
	},
	useCoupon: function(id){
		this.props.useCoupon(id);
	},
	render: function(){
		var DISPLAY;
		if(this.props.tickets == "Active Coupons Does Not Exist"){
			DISPLAY = <div>Active Coupons Does Not Exist</div>;
		}else{
			var COUPONS = this.state.tickets.map(function(coupon) {
            return <div className="row">
                        <div className="col-md-8">
                            <Coupon type={coupon.TYPE} 
										day={coupon.DAY} month={coupon.MONTH} year={coupon.YEAR}
										name={coupon.NAME}
										address={coupon.ADDRESS} township={coupon.TOWNSHIP} campus={coupon.CAMPUS}/>
                        </div>
                        <div clasName="col-md-4">
                            <button type="button"
                                className="btn btn-md btn-danger"
                                onClick={this.useCoupon.bind(this, coupon.COUPON_ID)}>
                                    Use
                            </button>
                        </div>
                    </div>
            ;
			}.bind(this));
			DISPLAY = <div>{COUPONS}</div>;
		}
		return(
			<div>
				{DISPLAY};
			</div>	
		);
	}
});

module.exports = Display;
