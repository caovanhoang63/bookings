const CSRFToken = document.currentScript.getAttribute('csrf_token');
const RoomID =  document.currentScript.getAttribute('room_id');
const checkAvailabilityButton = document.getElementById('check-availability-button');

checkAvailabilityButton.addEventListener('click', function () {
    let html= `
                <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
                    <div class="row">
                        <div class="col">
                            <div class="row" id="reservation-dates-modal">
                                <div class="col">
                                    <input disabled type="text" required class="form-control " id="start" name="start_date" placeholder="Arrival">
                                </div>
                                <div class="col">
                                    <input disabled  type="text" required class="form-control " id="end" name="end_date" placeholder="Departure">
                                </div>
                            </div>
                        </div>
                    </div>
                </form> `;

    Prompt().custom({
        msg: html,
        title: "Choose your date",
        willOpen: () => {
            //datepicker
            const elem = document.getElementById('reservation-dates-modal');
            const datepicker = new DateRangePicker(elem, {
                format: "dd-mm-yyyy",
                showOnFocus: true,
                orientation: 'top',
                minDate: new Date(),
            });
        },
        callback: function (result){
            let form = document.getElementById('check-availability-form');
            let formData = new FormData(form);
            formData.append("csrf_token", CSRFToken);
            formData.append("room_id",RoomID);
            fetch("/search-availability-json", {
                method: "POST",
                body: formData,
            })
                .then(response => response.json())
                .then(data => {
                    if (data.ok) {
                        attention.custom({
                            icon: 'success',
                            msg: `<p>Room is available!</p>`
                                + `<p><a href="/book-room?id=`
                                + data.room_id
                                + `&s=`
                                + data.start_date
                                + `&e=`
                                + data.end_date
                                +`"class="btn btn-primary">Book now!</a></p>`,
                            showConfirmButton: false,
                        });
                    } else {
                        attention.error({
                            icon: 'error',
                            msg: 'No availability',
                        });
                    }
                })
        }
    });
});