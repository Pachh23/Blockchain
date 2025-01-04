import React, { useEffect, useState } from 'react';
import { Form, Select, DatePicker, TimePicker, Input, Button, Card, Row, Col, Table, message } from 'antd';
import moment from 'moment';
import { RoomInterface } from '../../interface/IRoom';
import { DepartmentInterface } from '../../interface/IDepartment';
import { GetDepartment, GetRoom } from '../../services/https';

const { Option } = Select;
const { TextArea } = Input;

const AppointmentForm1: React.FC = () => {
  const [form] = Form.useForm();
  const [messageApi, contextHolder] = message.useMessage();
  const [selectedDept, setSelectedDept] = useState(null);
  const [rooms, setRooms] = useState<RoomInterface[]>([]);
  const [departments, setDepartments] = useState<DepartmentInterface[]>([]);

  
  const getDepartments = async () => {
    let res = await GetDepartment();
    if (res.status === 200) {
      setDepartments(res.data);
    } else {
      messageApi.error("Departments not found");
    }
  };
  
    // Fetch Initial Data
    const getRooms = async () => {
      let res = await GetRoom();
      if (res.status === 200) {
        setRooms(res.data);
      } else {
        messageApi.error("Rooms not found");
      }
    };


    useEffect(() => {
      getDepartments();
      getRooms();
    }, []);

    useEffect(() => {
      console.log(departments);
      console.log(rooms);
    }, [departments, rooms]);
    
    const handleDepartmentChange = (value: React.SetStateAction<null>) => {
      setSelectedDept(value);
      // Reset room selection when department changes
      form.setFieldValue('RoomID', undefined);
    };
  // กำหนดวันไม่ให้เลือกวันในอดีต
  const disabledDate = (current: moment.Moment) => {
    return current && current < moment().startOf('day');
  };

  // ข้อมูลสำหรับตาราง
  const columns = [
    {
      title: "Patient ID",
      dataIndex: "patient_id",
      key: "patient_id",
    },
    {
      title: "Doctor ID",
      dataIndex: "doctor_id",
      key: "doctor_id",
    },
    {
      title: "Date",
      dataIndex: "date",
      key: "date",
    },
    {
      title: "Time",
      dataIndex: "time",
      key: "time",
    },
    {
      title: "Reason",
      dataIndex: "reason",
      key: "reason",
    },
  ];

  return (
    <div style={{ padding: '24px' }}>
      {contextHolder}
      <Card title="Medical Appointment Form" bordered={false}>
        <Row gutter={16}>
          <Col xs={24} sm={12}>
            {/* ฟอร์มสำหรับเลือกแผนกและห้อง */}
            <Form
              form={form}
              layout="vertical"
              //onFinish={onFinish}
              autoComplete="off"
            >
              <Row gutter={16}>
              <Col xs={24} sm={12}>
          <Form.Item
            name="department_id"
            label="Department"
            rules={[{ required: true, message: 'Please select a department' }]}
          >
            <Select
              placeholder="Select department"
              onChange={handleDepartmentChange}
              loading={departments.length === 0}
            >
              {departments.map((item) => (
                <Select.Option value={item.ID} key={item.ID}>
                  {item.department}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
        </Col>
              <Col xs={24} sm={12}>
          <Form.Item
            name="RoomID"
            label="Examination Room"
            rules={[{ required: true, message: 'Please select a room' }]}
          >
            <Select
              placeholder="Select examination room"
              disabled={!selectedDept}
              loading={rooms.length === 0}
            >
              {rooms
                .filter((room) => room.department_id === selectedDept)
                .map((room) => (
                  <Select.Option key={room.ID} value={room.ID}>
                    {room.name}
                  </Select.Option>
                ))}
            </Select>
          </Form.Item>
        </Col>
                
              </Row>
              <Row gutter={16}>
                <Col xs={24} sm={12}>
                  <Form.Item
                    name="date"
                    label="Date"
                    rules={[{ required: true, message: 'Please select a date' }]}
                  >
                    <DatePicker 
                      style={{ width: '100%' }}
                      format="DD/MM/YYYY"
                      //disabled={disabledDate}
                    />
                  </Form.Item>
                </Col>

                <Col xs={24} sm={12}>
                  <Form.Item
                    name="time"
                    label="Time"
                    rules={[{ required: true, message: 'Please select a time' }]}
                  >
                    <TimePicker 
                      style={{ width: '100%' }}
                      format="HH:mm"
                      minuteStep={15}
                      showNow={false}
                    />
                  </Form.Item>
                </Col>
              </Row>

              <Form.Item
                name="symptoms"
                label="Initial Symptoms"
                rules={[{ required: true, message: 'Please describe your symptoms' }]}
              >
                <TextArea 
                  rows={4} 
                  placeholder="Please describe your initial symptoms"
                />
              </Form.Item>

              <Form.Item>
                <Button type="primary" htmlType="submit" block>
                  Confirm Appointment
                </Button>
              </Form.Item>
            </Form>
          </Col>

          {/* ตารางแสดงข้อมูลการนัดหมาย */}
          <Col xs={24} sm={12}>
            <Card title="Appointment Table" bordered={true}>
              <Table
                columns={columns}
                //dataSource={data}
                rowKey="patient_id"
              />
            </Card>
          </Col>
        </Row>
      </Card>
    </div>
  );
};

export default AppointmentForm1;
