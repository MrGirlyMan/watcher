var r = React.createElement;

var MainLayout = React.createClass({
	render: function() {
		return (
			r("div", {className: "app"},
				r("header", {className: "primary-header"}),
				r(Nav, null),
				r("main", null, this.props.children)
			)
		)
	}
})
