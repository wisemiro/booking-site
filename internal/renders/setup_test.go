package renders

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/models"
)

var session *scs.SessionManager

var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	//config
	testApp.InProduction = false //change to true in production, to change secure = true.

	//sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	testApp.Sessions = session

	app = &testApp
	
	os.Exit(m.Run())
}
