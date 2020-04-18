package htmlTemplates

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
)

func CreateNavbar() string {
	return fmt.Sprintf(NavbarTemplate, config.GetNamespaceOrDefault())
}

const NavbarTemplate = `
<nav class="navbar navbar-dark fixed-top bg-dark flex-md-nowrap p-0 shadow">
  <a class="navbar-brand col-sm-3 col-md-2 mr-0" href="#">OC-APM</a>
  <input class="form-control form-control-dark w-100" type="text" placeholder="Search" aria-label="Search">
  <ul class="navbar-nav px-3">
	<li class="nav-item text-nowrap">
	  <a class="nav-link" href="#">%s</a>
	</li>
  </ul>
</nav>
`

func CreateContentSideNavigation(sideLinks, mainContent string) string {
	return fmt.Sprintf(SideNavigationContentTemplate, sideLinks, mainContent)
}

const SideNavigationContentTemplate = `
<div class="container-fluid">
  <div class="row">
	<nav class="col-md-2 d-none d-md-block bg-light sidebar">
	  <div class="sidebar-sticky">
		<ul class="nav flex-column">
	      %[1]s
		</ul>
	   </div>
	</nav>

	%[2]s
  </div>
</div>
`

func CreateStaticNavigation() string {
	return fmt.Sprint(
		CreateNavItem("/", "home", "Home"),
		CreateNavItem(urls.GitPageUrl, "git", "Git"),
		CreateNavItemCommandsHeader(),
		CreateNavItem(urls.OcCmdList, "list", "List"),
		CreateNavItemAppsHeader(),
		CreateNavItem(urls.AppListPageUrl, "all", "All"),
		CreateAppsNavItems(),
	)
}

func CreateAppsNavItems() string {
	Apps, err := util.GetAppsFromDisk()
	if err != nil {
		return err.Error()
	}

	items := ""
	for _, p := range Apps {
		items = items + CreateNavItem(urls.AppDetailsPageUrl+p.Name, p.Name, "# "+p.Name)
	}
	return items
}

func CreateNavItem(href, id, title string) string {
	return fmt.Sprintf(NavItemTemplate, href, id, title)
}

const NavItemTemplate = `<li class="nav-item">
   <a class="nav-link" href="%[1]s">
   %[3]s
   </a>
</li>`

func CreateNavItemAppsHeader() string {
	return fmt.Sprintf(NavItemAppsHeader, urls.AppNewPageUrl)
}

const NavItemAppsHeader = `<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted">
  <span>Applications</span>
  <a class="d-flex align-items-center text-muted" href="%s" aria-label="Add a new App">
	<b>+ New</b>
  </a>
</h6>`

func CreateNavItemCommandsHeader() string {
	return fmt.Sprintf(NavItemCommandsHeader)
}

const NavItemCommandsHeader = `<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted">
  <span>OC-Commands</span>
</h6>`

func CreateMainContent(title string, content string) string {
	return fmt.Sprintf(MainContentTemplate, title, content)
}

const MainContentTemplate = `
<main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-4">
	<br />
	<h1>%[1]s</h1>
	<hr />
	%[2]s
</main>
`
