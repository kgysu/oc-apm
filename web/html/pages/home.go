package htmlPages

import (
	"fmt"
	"github.com/kgysu/oc-apm/config"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
)

// Home/Main Page
func CreateHomePage() string {
	return htmlTemplates.CreateHtml("Home", fmt.Sprint(
		htmlTemplates.CreateBreaks(2),
		fmt.Sprintf(homePage, config.GetVersion(), config.HelpText),
		htmlTemplates.CreateBreaks(5),
	))
}

const homePage = `
<div class="jumbotron bg-light">
  <h1 class="display-4">Welcome!</h1>
  <p class="lead">OC-APM is a helpful tool to manage applications on an Openshift-Cluster.</p>
  <hr class="my-4">
  <p>Start with creating a new app.</p>
  <a class="btn btn-primary btn-lg" href="/new/app" role="button">Create first app</a>
</div>
<br />
<br />
<br />
<br />
<br />
<hr />
<br />
<h2>Help</h2>
<br />
<h4>Info</h4>
<br />
<p>GitHub-Repository: <a class="" href="https://github.com/kgysu/oc-apm" target="_blank">github.com/kgysu/oc-apm</a></p>
<p><b>Version:</b> %s</p>
<br />
<h5>Run APM local</h5>
<p>If you want to run apm on your local machine then ensure to login with oc client first:</p>
<kbd>oc login</kbd> <br />
<kbd>oc project default</kbd> <br />
<br />
<br />
<h4>App Management</h4>
<br />
<h5>File structure</h5>
<p>Manage your apps locally in a file structure like this:</p>
<div class="yaml-content">
<pre><code>
/apps/
  /app-one/
    - app-DeploymentConfig.yaml
    - app-Service.yaml
    - ...yaml
    - /environment-dev/
      - dev.env
    - /environment-prod/
      - prod.env

  /app-two
    - app-DeploymentConfig.yaml
    - app-Service.yaml
    - ...yaml
</pre></code>
</div>
<br />
<br />
<h5>Yaml files</h5>
<p>Yaml-Files under <span class="badge badge-secondary">/apps</span> are basically Openshift object(kind) descriptions. 
The following kinds are supported:</p>
<ul>
	<li>DeploymentConfig</li>
	<li>StatefulSet</li>
	<li>Service</li>
	<li>Route</li>
	<li>ConfigMap</li>
	<li>PersistentVolumeClaim</li>
	<li>ServiceAccount</li>
	<li>Role</li>
	<li>RoleBinding</li>
</ul>
<br />
<br />
<h5>Env files</h5>
<p>Env-Files are used to set namespace specific values within yaml-files. You can set a variable in a yaml file like 
this <span class="badge badge-secondary">${name}</span></p>
<p>For example:</p>
<div class="yaml-content">
<pre><code>
apiVersion: apps.openshift.io/v1
kind: DeploymentConfig
metadata:	
...
  - env:
    - name: SAMPLE
      value: ${VARIABLE}
    - name: TWO
      value: ${VARIABLE_TWO}
...
</pre></code>
</div>
<br />
<p>And then define the values in the env-files like this:</p>
<div class="yaml-content">
	<pre><code>
VARIABLE=some-value
VARIABLE_TWO=some-other-value
...
	</pre></code>
</div>
<br />
<br />
<br />
<h4>command-line help</h4>
<div class="yaml-content">
	<pre><code>
	%s
	</pre></code>
</div>
<br />
<br />
<br />
`
