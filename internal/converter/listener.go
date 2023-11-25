package xdsconverter

import (
	"strings"
	xdsmodels "xds_server/internal/models"
)

const (
	listenerAddress = "127.0.0.1"
	safeMode        = "safe"
	unsafeMode      = "unsafe"

	listener = "listener"
)

func ListenerConverter(domainEp *xdsmodels.DomainEndpoint, cert *xdsmodels.Cert, port uint32) *xdsmodels.Listener {
	var routeName strings.Builder
	routeName.WriteString(route)
	routeName.WriteString("-")
	routeName.WriteString(domainEp.Domain)

	var genListener xdsmodels.Listener

	genListener.Address = listenerAddress
	genListener.RouteConfigName = routeName.String()
	genListener.Port = port

	if cert == nil {
		//UNSAFE LISTENER
		var listenerUnsafeName strings.Builder
		listenerUnsafeName.WriteString(unsafeMode)
		listenerUnsafeName.WriteString(listener)
		listenerUnsafeName.WriteString(domainEp.Domain)

		listenerUnsafe := genListener
		listenerUnsafe.Name = listenerUnsafeName.String()

		return &listenerUnsafe
	}

	//SAFE LISTENER
	var listenerSafeName strings.Builder
	listenerSafeName.WriteString(safeMode)
	listenerSafeName.WriteString(listener)
	listenerSafeName.WriteString(domainEp.Domain)

	listenerSafe := genListener
	listenerSafe.Name = listenerSafeName.String()
	listenerSafe.Cert = &xdsmodels.Cert{
		CertificateChain: cert.CertificateChain,
		PrivateKey:       cert.PrivateKey,
	}

	return &listenerSafe
}
