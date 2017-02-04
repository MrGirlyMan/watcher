var h = React.createElement;

var HelloWorldBanner =  React.createClass({
	render: function(){
		return (
			h('div', {className: "container"},
				h('h1', null, 'Hello World'),
				h('div', {className: "starter-template"},
					h('p', {className: 'lead'}, 'Testing testing testing...' )
				)
			)
		)
	}
})
