class Phone:
    def __init__(self, model, battery, f_memory, emkost, volume):
        self.model = model
        self.battery = battery
        self.f_memory = f_memory
        self.emkost = emkost
        self.volume = volume

    def get_number(self, r_number):
        contacts = {
            "+79345462739": "Степан Андреевич",
            "+79145462738": "Татьяна Николаевна",
            "+79345462738": "Юлия Яблокова",
            "+79355462738": "Иван Иванович",
            "+79340462738": "Дмитрий Степанович",
            "+79345468738": "Степан Дмитриевич"
        }
        return contacts.get(r_number, "Введите номер")

    def time_battery(self, power):
        ost = (100 - self.battery) * self.emkost / 100
        skorost = (power / 5) * 1000
        if skorost == 0:
            return 0
        tm_bat = ost / skorost * 60
        return tm_bat

    def new_volume(self, nazhatv, nazhatn):
        kvol = self.volume + nazhatv - nazhatn
        return kvol

def ready_labaa6():
    phone = Phone("Xiaomi13", 89, 189, 5000, 2)
    print("Вам звонит", phone.get_number("+79345468738"))
    print("У вас", phone.battery, "%")
    print("Чтобы зарядить батарею вам нужно {:.2f} минут.".format(phone.time_battery(5)))
    print("Звук установлен с", phone.volume, "на", phone.new_volume(2, 1))

ready_labaa6()