var React = require('react');

var Window = React.createClass({
    render: function(){
        return (
            <div>
                <p>Testing React Window</p>
            </div>
        );
    }
});

React.render(<Window/>, document.getElementById('app'));
