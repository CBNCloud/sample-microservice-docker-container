package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
//"github.com/gin-gonic/gin"
)


func TestLoginFormPass(t *testing.T) {
        passnya := "123"

        if passnya != "456" {
                assert := assert.New(t)
                assert.NotEqual(passnya, false, "seharusnya 123")
        }
}
