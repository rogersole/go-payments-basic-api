package apis

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/satori/go.uuid"
)

// paymentService specifies the interface for the payment service needed by serviceResource
type paymentService interface {
	Get(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error)
	Query(rs app.RequestScope, offset, limit int) ([]dtos.Payment, error)
	Count(rs app.RequestScope) (int, error)
	Create(rs app.RequestScope, model *dtos.Payment) (*dtos.Payment, error)
	Update(rs app.RequestScope, id uuid.UUID, model *dtos.Payment) (*dtos.Payment, error)
	Delete(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error)
}

// paymentResource defines the handlers for the CRUD APIs
type paymentResource struct {
	service paymentService
}

// ServePayment sets up the routing of payment endpoints and the corresponding handlers
func ServePaymentResource(rg *routing.RouteGroup, service paymentService) {
	r := &paymentResource{service}
	rg.Get("/payments/<id>", r.get)
	rg.Get("/payments", r.query)
	rg.Post("/payments", r.create)
	rg.Put("/payments/<id>", r.update)
	rg.Delete("/payments/<id>", r.delete)
}

func (r *paymentResource) get(c *routing.Context) error {
	id := uuid.Must(uuid.FromString(c.Param("id")))
	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *paymentResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *paymentResource) create(c *routing.Context) error {
	var model dtos.Payment
	if err := c.Read(&model); err != nil {
		return err
	}
	// TODO: remove following line!!
	spew.Dump(model)
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *paymentResource) update(c *routing.Context) error {
	id := uuid.Must(uuid.FromString(c.Param("id")))
	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *paymentResource) delete(c *routing.Context) error {
	id := uuid.Must(uuid.FromString(c.Param("id")))
	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
