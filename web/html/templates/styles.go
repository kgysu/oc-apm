package htmlTemplates

func CreateStyles() string {
	return MainStyles
}

const MainStyles = `
<style>
	body {
	  font-size: 1rem;
	}
	
	.feather {
	  width: 16px;
	  height: 16px;
	  vertical-align: text-bottom;
	}
	
	/*
	 * Sidebar
	 */
	.sidebar {
	  position: fixed;
	  top: 0;
	  bottom: 0;
	  left: 0;
	  z-index: 100; /* Behind the navbar */
	  padding: 48px 0 0; /* Height of navbar */
	  box-shadow: inset -1px 0 0 rgba(0, 0, 0, .1);
	}
	
	.sidebar-sticky {
	  position: relative;
	  top: 0;
	  height: calc(100vh - 48px);
	  padding-top: .5rem;
	  overflow-x: hidden;
	  overflow-y: auto; /* Scrollable contents if viewport is shorter than content. */
	}
	
	@supports ((position: -webkit-sticky) or (position: sticky)) {
	  .sidebar-sticky {
		position: -webkit-sticky;
		position: sticky;
	  }
	}
	
	.sidebar .nav-link {
	  font-weight: 500;
	  color: #333;
	}
	
	.sidebar .nav-link .feather {
	  margin-right: 4px;
	  color: #999;
	}
	
	.sidebar .nav-link.active {
	  color: #007bff;
	}
	
	.sidebar .nav-link:hover .feather,
	.sidebar .nav-link.active .feather {
	  color: inherit;
	}
	
	.sidebar-heading {
	  font-size: .75rem;
	  text-transform: uppercase;
	}
	
	/*
	 * Content
	 */
	[role="main"] {
	  padding-top: 133px; /* Space for fixed navbar */
	}
	
	@media (min-width: 768px) {
	  [role="main"] {
		padding-top: 48px; /* Space for fixed navbar */
	  }
	}
	
	/*
	 * Navbar
	 */
	.navbar-brand {
	  padding-top: .75rem;
	  padding-bottom: .75rem;
	  font-size: 1rem;
	  background-color: rgba(0, 0, 0, .25);
	  box-shadow: inset -1px 0 0 rgba(0, 0, 0, .25);
	}
	
	.navbar .form-control {
	  padding: .75rem 1rem;
	  border-width: 0;
	  border-radius: 0;
	}
	
	.form-control-dark {
	  color: #fff;
	  background-color: rgba(255, 255, 255, .1);
	  border-color: rgba(255, 255, 255, .1);
	}
	
	.form-control-dark:focus {
	  border-color: transparent;
	  box-shadow: 0 0 0 3px rgba(255, 255, 255, .25);
	}

	.yaml-content {
      background: #f8f9fa;
      border: solid 24px #f8f9fa;
      box-shadow: 4px 4px 20px 1px rgba(0, 0, 0, .1);
	}

    .table-shadow {
      box-shadow: 1px 1px 20px 1px rgba(0, 0, 0, .1);
	}
	

</style>
`