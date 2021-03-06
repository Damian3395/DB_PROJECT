var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var request = require('request');
var Options = require('../../misc/Options.jsx');
var Coupon = require('./Coupon.jsx');
var Error = require('../../misc/Error.jsx')

var Search = React.createClass({
	getInitialState: function(){
		return{error: "", tickets: []};	
	},
	createTicket: function(id){
		request({
			url: 'http://www.ruexploring.com:80/CreateTicket',
			method: 'POST',
			json: {
				ID: this.props.userID,
				COUPON_ID: id
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
				if(body != "Max Active Tickets Reached"){
					$('#query').val("");
					$('#min_age').val("0");
					$('#max_age').val("0");
					$('#min_people').val("0");
					$('#max_people').val("0");
					this.setState({tickets: []});
				}else{
					this.setState({error: body, tickets: []});	
				}
			}
		}.bind(this));
	},
	getSearchResult: function(){
		request({
			url: 'http://www.ruexploring.com:80/QueryCoupon',
			method: 'POST',
			json: {
				ID: this.props.userID,
				QUERY: $('#query').val(),
				MAIN_CATEGORY: $('#main_category').val(),
				SUB_CATEGORY: $('#sub_category').val(),
				MIN_AGE: $('#min_age').val(),
				MAX_AGE: $('#max_age').val(),
				MIN_PEOPLE: $('#min_people').val(),
				MAX_PEOPLE: $('#max_people').val(),
				CAMPUS: $('#campus').val()
			}
		}, function(error, response, body){
			if(error){
				console.log(error);
			}else{
				console.log(response.statusCode, body);
				var objects = _.cloneDeep(body.coupons);
                this.setState({tickets: objects});
			}
		}.bind(this));
		this.forceUpdate();
	},
	render: function(){
		var COUPONS = this.state.tickets.map(function(coupon) {
            return <div className="row">
                        <div className="col-md-8">
                            <Coupon type={coupon.TYPE} day={coupon.DAY} month={coupon.MONTH} year={coupon.YEAR}
									name={coupon.NAME} address={coupon.ADDRESS} township={coupon.TOWNSHIP} campus={coupon.CAMPUS}/>
                        </div>
                        <div clasName="col-md-4">
                            <button type="button"
                                className="btn btn-md btn-danger"
                                onClick={this.createTicket.bind(this, coupon.COUPON_ID)}>
                                    Add
                            </button>
                        </div>
                    </div>
            ;
        }.bind(this));
        var DISPLAY = <div>{COUPONS}</div>;
	
		return (	
			<div>
				<div className="row">
					<div className="col-md-12">
						<Error error={this.state.error}/>
						<h3 className="text-center"><strong>Search For Coupons</strong></h3>
					</div>
				</div>
				<div className="row">
					<div className="col-md-12">
						<div className="input-group">
							<span className="input-group-addon">Search For Key Words:</span>
							<input type="text" className="form-control" id="query"/>
						</div>
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Main Category: </span>
							<select	className="form-control" id="main_category">
								<Options type="main"/>
							</select>
						</div>	
					</div>
				</div>	
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Sub Category: </span>
							<select	className="form-control" id="sub_category">
								<Options type="sub"/>
							</select>
						</div>	
					</div>
				</div>	
				<br/>
				<div className="row">
					<div className="col-md-8">
						<div className="input-group">
							<span className="input-group-addon">Campus: </span>
							<select	className="form-control" id="campus">
								<option value="New Brunswick">New Brunswick</option>
								<option value="Newark">Newark</option>
								<option value="Camden">Camden</option>
							</select>
						</div>	
					</div>
				</div>	
				<br/>
				<div className="row">
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Min Age:</span>
							<input type="number" className="form-control" id="min_age"/>
						</div>	
					</div>
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Max Age:</span>
							<input type="number" className="form-control" id="max_age"/>
						</div>	
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Min People:</span>
							<input type="number" className="form-control" id="min_people"/>
						</div>	
					</div>
					<div className="col-md-4">
						<div className="input-group">
							<span className="input-group-addon">Max People:</span>
							<input type="number" className="form-control" id="max_people"/>
						</div>	
					</div>
				</div>
				<br/>
				<div className="row">
					<div className="col-md-offset-8 col-md-4">
						<button type="button"
							className="btn btn-lg btn-block btn-primary"
							onClick={this.getSearchResult}>
								Search
						</button>
					</div>	
				</div>
				<br/>
				<div className="row">
					<div className="col-md-12">
						<h3 className="text-center"><strong>Coupons Found</strong></h3>
					</div>
				</div>
				{DISPLAY}
			</div>
		);
	}
});

module.exports = Search;
