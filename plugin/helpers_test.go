package plugin_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/mattermost/mattermost-server/plugin/plugintest"
	"github.com/mattermost/mattermost-server/plugin/plugintest/mock"
	"github.com/mattermost/mattermost-server/utils/fileutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstallPluginFromUrl(t *testing.T) {
	replace := true
	h := &plugin.HelpersImpl{}
	api := &plugintest.API{}
	h.API = api

	t.Run("error while parsing the download url", func(t *testing.T) {
		_, err := h.InstallPluginFromUrl("http://%41:8080/", replace)

		assert.NotNil(t, err)
		assert.Equal(t, "error while parsing url: parse http://%41:8080/: invalid URL escape \"%41\"", err.Error())
	})

	t.Run("errors out while downloading file", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusInternalServerError)
		}))
		defer testServer.Close()
		url := testServer.URL

		_, err := h.InstallPluginFromUrl(url, replace)

		assert.NotNil(t, err)
		assert.Equal(t, "received 500 status code while downloading plugin from server", err.Error())
	})

	t.Run("downloads the file successfully", func(t *testing.T) {
		path, _ := fileutils.FindDir("tests")
		tarData, err := ioutil.ReadFile(filepath.Join(path, "testplugin.tar.gz"))
		require.NoError(t, err)
		expectedManifest := &model.Manifest{Id: "testplugin"}
		api.On("InstallPlugin", mock.Anything, false).Return(expectedManifest, nil)

		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
			_, _ = res.Write(tarData)
		}))
		defer testServer.Close()
		url := testServer.URL

		manifest, err := h.InstallPluginFromUrl(url, false)

		assert.Nil(t, err)
		assert.Equal(t, "testplugin", manifest.Id)
	})

	t.Run("the url pointing to server is incorrect", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusNotFound)
		}))
		defer testServer.Close()
		url := testServer.URL

		_, err := h.InstallPluginFromUrl(url, false)

		assert.NotNil(t, err)
		assert.Equal(t, "received 404 status code while downloading plugin from server", err.Error())
	})
}
