var r = React.createElement;

var Nav = React.createClass({
	render: function() {
		return (
			r("nav", {className: "navbar navbar-inverse navbar-fixed-top"},
				r("div", {className: "container-fluid"},
					r("div", {className: "navbar-header"},
						r("button", {className: "navbar-toggle collapsed", "data-toggle": "collapse", "data-target": "navbar"},
							r("span", {className: "sr-only"}, "Toggle Navigation"),
							r("span", {className: "icon-bar"}),
							r("span", {className: "icon-bar"}),
							r("span", {className: "icon-bar"})
						),
						r("a", {className: "navbar-brand", href: "/"}, "Watcher")
					),
					r("div", {className: "navbar-collapse collapse", id: "navbar"},
						r("ul", {className: "nav navbar-nav navbar-right"},
							r("li", null,
								r("a", {href: "/"}, "Account")
							),
							r("li", null,
								r("a", {href: "/"}, "Login")
							)
						)
					)
				)
			)
		)
	}
})
