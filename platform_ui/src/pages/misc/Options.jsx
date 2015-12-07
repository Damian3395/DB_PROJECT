var React = require('react');

var Options = React.createClass({
	render: function(){
		if(this.props.type == "main"){
			return (
				<div>
					<option value="Food">Food</option>
					<option value="Cafe">Cafe</option>
					<option value="Sport">Sport</option>
					<option value="Outdoor">Outdoor</option>
				</div>
			);
		}else if(this.props.type == "sub"){
			return (
				<div>
					<option value="Fast food">Fast food</option>
					<option value="Mediterranean">Mediterranean</option>
					<option value="Chinese">Chinese</option>
					<option value="Indian">Indian</option>
					<option value="Japanese">Japanese</option>
					<option value="Mexican">Mexican</option>
					<option value="Vietnamese">Vietnamese</option>
					<option value="Italian">Italian</option>
					<option value="Spanish">Spanish</option>
					<option value="South American">South American</option>
					<option value="Thai">Thai</option>
					<option value="Romanian">Romanian</option>
					<option value="Southern">Southern</option>
					<option value="Diner">Diner</option>
					<option value="BBQ">BBQ</option>
					<option value="Polish">Polish</option>
					<option value="Bar">Bar</option>
					<option value="Bubble Tea">Bubble Tea</option>
					<option value="Coffee Shop">Coffee Shop</option>
					<option value="Bowling">Bowling</option>
					<option value="Golf">Golf</option>
					<option value="Archery">Archery</option>
					<option value="Shooting Range">Shooting Range</option>
					<option value="Go Karting">Go Karting</option>
					<option value="Swimming">Swimming</option>
					<option value="Surfing">Surfing</option>
					<option value="Skiing">Skiing</option>
					<option value="Camping">Camping</option>
					<option value="Mountain Biking">Mountain Biking</option>
					<option value="Canoeing">Canoeing</option>
					<option value="Sailing">Sailing</option>
				</div>
			);
		}
	}
});

module.exports = Options;
